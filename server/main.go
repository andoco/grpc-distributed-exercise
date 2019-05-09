package main

import (
	"flag"
	"log"

	"github.com/andoco/ably-distributed-exercise/server/stateless"
)

func main() {
	port := flag.Int("port", 8888, "")
	flag.Parse()

	cmd := flag.Arg(0)

	switch cmd {
	case "stateless":
		if err := stateless.Serve(*port); err != nil {
			log.Fatal(err)
		}
	case "stateful":
		log.Println("would run stateful")
	default:
		log.Println("USAGE: [OPTS] <stateless|stateful>")
	}
}
