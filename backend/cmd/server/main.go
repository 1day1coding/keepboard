package main

import (
	"log"

	"keepboard/backend/pkg/server"
)

func main() {
	s := server.New(server.Option{Port: "8888"})
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
