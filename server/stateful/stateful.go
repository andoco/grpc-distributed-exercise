package stateful

import (
	"fmt"
	"net"

	"github.com/andoco/ably-distributed-exercise/server/stateful/randstream"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type service struct{}

func (s *service) Begin(req *randstream.BeginRequest, stream randstream.Generator_BeginServer) error {
	stream.Send(&randstream.Number{Value: 1})
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
