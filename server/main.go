package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/andoco/ably-distributed-exercise/server/numbers"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type service struct{}

func (s *service) Begin(req *numbers.BeginRequest, stream numbers.Generator_BeginServer) error {
	var a int32 = 1
	for {
		time.Sleep(1 * time.Second)
		if err := stream.Send(&numbers.Number{Value: a}); err != nil {
			return errors.Wrap(err, "sending number to stream")
		}
		a *= 2
	}
}

func serve(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	grpcServer := grpc.NewServer()
	numbers.RegisterGeneratorServer(grpcServer, &service{})

	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "serve")
	}

	return nil
}

func main() {
	port := flag.Int("port", 8888, "")
	flag.Parse()

	if err := serve(*port); err != nil {
		log.Fatal(err)
	}
}
