package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/endpoint"
	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/service"
	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/transport"
	"github.com/muhammadnurbasari/onesmile-test-protobuffer/proto/generate"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	logkit "github.com/go-kit/kit/log"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	PORT := os.Getenv("PORT")

	var logger logkit.Logger
	{
		logger = logkit.NewLogfmtLogger(os.Stderr)
		logger = logkit.With(logger, "ts", logkit.DefaultTimestampUTC)
		logger = logkit.With(logger, "caller", logkit.DefaultCaller)
	}

	var (
		service    = service.NewServiceCC()
		endpoints  = endpoint.NewEndpointCC(service)
		grpcServer = transport.NewGrpcServer(endpoints, logger)
	)
	srv := grpc.NewServer()
	generate.RegisterValidationServer(srv, grpcServer)

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
