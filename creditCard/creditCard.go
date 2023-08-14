package creditCard

import (
	"context"
	"errors"
	"strconv"

	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/luhn"
	"github.com/muhammadnurbasari/onesmile-test-protobuffer/proto/generate"
	"github.com/rs/zerolog/log"
)

type CreditCard struct{}

func (c CreditCard) ValidateCreditCard(ctx context.Context, param *generate.CreditCard) (*generate.Validate, error) {
	number, err := strconv.Atoi(param.CreditCard)

	if err != nil {
		log.Error().Msg("error: " + err.Error())
		return nil, err
	}

	isValidate := luhn.Valid(number)

	if !isValidate {
		log.Error().Msg("error: " + errors.New("not a valid credit card").Error())
		var response = generate.Validate{IsValidate: isValidate}
		return &response, nil
	}

	log.Info().Msg("success : credit card is valid")
	var response = generate.Validate{IsValidate: isValidate}

	return &response, nil
}
