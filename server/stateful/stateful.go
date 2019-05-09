package stateful

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/andoco/ably-distributed-exercise/server/stateful/randstream"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type service struct{}

func (s *service) Begin(req *randstream.BeginRequest, stream randstream.Generator_BeginServer) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := int32(0); i < req.MaxNumbers; i++ {
		time.Sleep(1 * time.Second)
		a := r.Uint32()
		if err := stream.Send(&randstream.Number{Value: a}); err != nil {
			return errors.Wrap(err, "sending number to stream")
		}
	}

	if err := stream.Send(&randstream.Number{Checksum: "abcd1234"}); err != nil {
		return errors.Wrap(err, "sending number to stream")
	}

	return nil
}

func Serve(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	grpcServer := grpc.NewServer()
	randstream.RegisterGeneratorServer(grpcServer, &service{})

	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "serve")
	}

	return nil
}
