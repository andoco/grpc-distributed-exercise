package stateful

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/andoco/ably-distributed-exercise/server/stateful/randstream"
	"github.com/andoco/ably-distributed-exercise/server/stateful/session"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type service struct {
	sessionStore session.Store
}

func (s *service) generate(state *session.State, r *rand.Rand, stream randstream.Generator_BeginServer) error {
	for i := state.NumGenerated; i < state.MaxNumbers; i++ {
		time.Sleep(1 * time.Second)
		a := r.Uint32()

		reply := &randstream.Number{Value: a}
		if i == state.MaxNumbers-1 {
			reply.Checksum = "123456789"
		}

		state.LastActive = time.Now()

		if err := stream.Send(reply); err != nil {
			if err := s.sessionStore.Update(state); err != nil {
				return errors.Wrap(err, "could not update session")
			}

			return errors.Wrap(err, "sending number to stream")
		}

		state.NumGenerated++
		if err := s.sessionStore.Update(state); err != nil {
			return errors.Wrap(err, "could not update session")
		}
	}

	if err := s.sessionStore.Delete(state.ClientId); err != nil {
		return errors.Wrap(err, "could not delete session state")
	}

	return nil
}

func newRand(seed int64, from int) *rand.Rand {
	r := rand.New(rand.NewSource(seed))

	// Fast-forward the random numbers.
	for i := 0; i < from; i++ {
		r.Uint32()
	}

	return r
}

func (s *service) Begin(req *randstream.BeginRequest, stream randstream.Generator_BeginServer) error {
	state := &session.State{
		ClientId:   req.ClientId,
		MaxNumbers: int(req.MaxNumbers),
		Seed:       time.Now().UnixNano(),
		LastActive: time.Now(),
	}
	s.sessionStore.Add(state)

	r := newRand(state.Seed, 0)

	if err := s.generate(state, r, stream); err != nil {
		return errors.Wrap(err, "failed while generating value stream")
	}

	return nil
}

func (s *service) Resume(req *randstream.ResumeRequest, stream randstream.Generator_ResumeServer) error {
	state, err := s.sessionStore.Get(req.ClientId)
	if err != nil {
		return errors.Wrap(err, "could not resume")
	}

	r := newRand(state.Seed, state.NumGenerated)

	log.Printf("Resuming with state %+v", state)

	if err := s.generate(state, r, stream); err != nil {
		return errors.Wrap(err, "failed while generating value stream")
	}

	return nil
}

func Serve(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	errHandler := func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, ss)
		if err != nil {
			log.Printf("method %q failed: %s", info.FullMethod, err)
		}
		return err
	}

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(errHandler))
	service := &service{sessionStore: session.NewAStore()}
	randstream.RegisterGeneratorServer(grpcServer, service)

	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "serve")
	}

	return nil
}
