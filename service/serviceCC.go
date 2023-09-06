package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/muhammadnurbasari/onesmile-test-grpc-creditcard/luhn"
	"github.com/rs/zerolog/log"
)

type ServiceCC interface {
	ValidateCreditCard(ctx context.Context, creditCard string) (bool, error)
}

func NewServiceCC() ServiceCC {
	return &basicServiceCC{}
}

type basicServiceCC struct{}

func (b *basicServiceCC) ValidateCreditCard(_ context.Context, creditCard string) (bool, error) {
	numberCC, err := strconv.Atoi(creditCard)

	if err != nil {
		log.Error().Msg("error : " + "credit card must be number of string")
		return false, errors.New("credit card must be number of string")
	}

	isValidate := luhn.Valid(numberCC)

	if !isValidate {
		log.Error().Msg("error: " + errors.New("not a valid credit card").Error())
		return isValidate, errors.New("not a valid credit card")
	}

	log.Info().Msg("success : credit card is valid")

	return isValidate, nil
}
