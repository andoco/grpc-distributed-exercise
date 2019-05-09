package main

import (
	"flag"
	"log"

	"github.com/andoco/ably-distributed-exercise/server/stateless"
)

func main() {
	port := flag.Int("port", 8888, "")
	flag.Parse()

	if err := stateless.Serve(*port); err != nil {
		log.Fatal(err)
	}
}
