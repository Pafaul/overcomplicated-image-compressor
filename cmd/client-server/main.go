package main

import (
	"log"
	"pafaul/overcomplicated-image-compressor/client-server"
)

func main() {
	s := client_server.NewServer("localhost:8000")
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
