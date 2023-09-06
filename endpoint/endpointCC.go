package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/models"
	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/service"
)

type EndpointCC struct {
	ValidateCCEndpoint endpoint.Endpoint
}

func NewEndpointCC(svc service.ServiceCC) EndpointCC {
	var validateCCEndpoint endpoint.Endpoint
	{
		validateCCEndpoint = MakeEndpointCC(svc)
	}

	return EndpointCC{ValidateCCEndpoint: validateCCEndpoint}
}

func MakeEndpointCC(svc service.ServiceCC) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.CcRequest)
		isValidate, err := svc.ValidateCreditCard(ctx, req.CreditCard)

		return models.CcResponse{Isvalidate: isValidate}, err
	}
}
