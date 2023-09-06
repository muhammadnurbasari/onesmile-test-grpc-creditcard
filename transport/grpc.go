package transport

import (
	"context"

	logkit "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endp "github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/endpoint"
	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/models"
	"github.com/muhammadnurbasari/onesmile-test-protobuffer/proto/generate"
)

type grpcServer struct {
	validateCreditCard grpctransport.Handler
}

func NewGrpcServer(endpoints endp.EndpointCC, logger logkit.Logger) generate.ValidationServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	return &grpcServer{
		validateCreditCard: grpctransport.NewServer(
			endpoints.ValidateCCEndpoint,
			decodeGrpcCreditCardRequest,
			encodeGrpcCreditCardResponse,
			options...,
		),
	}
}

func (s *grpcServer) ValidateCreditCard(ctx context.Context, req *generate.CreditCard) (*generate.Validate, error) {
	_, resp, err := s.validateCreditCard.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp.(*generate.Validate), nil
}

func decodeGrpcCreditCardRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*generate.CreditCard)

	return models.CcRequest{CreditCard: req.CreditCard}, nil
}

func encodeGrpcCreditCardResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(models.CcResponse)
	return &generate.Validate{IsValidate: resp.Isvalidate}, nil
}
