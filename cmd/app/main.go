package main

import (
	"context"
	"log"

	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/http/rest"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	server, err := rest.NewServer()

	if err != nil {
		return err
	}

	err = server.Run(ctx)

	return err
}
