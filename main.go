package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/creditCard"
	"github.com/muhammadnurbasari/onesmile-test-protobuffer/proto/generate"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	PORT := os.Getenv("PORT")

	srv := grpc.NewServer()
	var cc creditCard.CreditCard
	generate.RegisterValidationServer(srv, cc)

	log.Info().Msg("Starting RPC server at " + ":" + PORT)

	l, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Error().Msg("could not listen to " + PORT + " error: " + err.Error())
		os.Exit(1)
	}

	if err := srv.Serve(l); err != nil {
		log.Error().Msg("could not listen to " + PORT + " error: " + err.Error())
		os.Exit(1)
	}
}
