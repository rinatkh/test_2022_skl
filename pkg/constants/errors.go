package constants

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type CodedError struct {
	err  error
	code int
}

func (ce *CodedError) Error() string {
	return ce.err.Error()
}

func (ce *CodedError) Code() int {
	return ce.code
}

func NewCodedError(errMessage string, code int) *CodedError {
	return &CodedError{errors.New(errMessage), code}
}

var (
	// Unathorized
	InputError            = &CodedError{errors.New("bad json request"), fiber.StatusBadRequest}
	ErrUserDBNotFound     = &CodedError{errors.New("user not found in the database"), fiber.StatusBadRequest}
	ErrProductDBNotFound  = &CodedError{errors.New("product not found in the database"), fiber.StatusBadRequest}
	ErrCurrencyDBNotFound = &CodedError{errors.New("currency not found in the database"), fiber.StatusBadRequest}
	ErrOrderDBNotFound    = &CodedError{errors.New("order not found in the database"), fiber.StatusBadRequest}
	AuthError             = &CodedError{errors.New("Invalid public api key"), fiber.StatusUnauthorized}
	ErrConvertData        = &CodedError{errors.New("failed to convert"), fiber.StatusInternalServerError}
	ErrDB                 = &CodedError{errors.New("failed working with db"), fiber.StatusInternalServerError}
)
