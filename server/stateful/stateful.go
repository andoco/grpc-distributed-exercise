package stateful

import (
	"fmt"
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

func (s *service) generate(state *session.State, stream randstream.Generator_BeginServer) error {
	for i := state.NumGenerated; i < state.MaxNumbers; i++ {
		time.Sleep(1 * time.Second)
		a := state.Rand.Uint32()

		if err := stream.Send(&randstream.Number{Value: a}); err != nil {
			return errors.Wrap(err, "sending number to stream")
		}

		state.NumGenerated++
		s.sessionStore.Update(state)
	}

	if err := stream.Send(&randstream.Number{Checksum: "abcd1234"}); err != nil {
		return errors.Wrap(err, "sending number to stream")
	}

	return nil
}

func (s *service) Begin(req *randstream.BeginRequest, stream randstream.Generator_BeginServer) error {
	state := &session.State{
		ClientId:   req.ClientId,
		Rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
		MaxNumbers: int(req.MaxNumbers),
	}
	s.sessionStore.Add(state)

	if err := s.generate(state, stream); err != nil {
		return errors.Wrap(err, "failed while generating value stream")
	}

	return nil
}

func (s *service) Resume(req *randstream.ResumeRequest, stream randstream.Generator_ResumeServer) error {
	state, err := s.sessionStore.Get(req.ClientId)
	if err != nil {
		return errors.Wrap(err, "could not resume")
	}

	if err := s.generate(state, stream); err != nil {
		return errors.Wrap(err, "failed while generating value stream")
	}

	return nil
}

func Serve(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	grpcServer := grpc.NewServer()
	service := &service{sessionStore: session.NewAStore()}
	randstream.RegisterGeneratorServer(grpcServer, service)

	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "serve")
	}

	return nil
}
