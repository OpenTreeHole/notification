package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HttpError struct {
	Code    int          `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
	Detail  *ErrorDetail `json:"detail,omitempty"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func MyErrorHandler(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	httpError := HttpError{
		Code:    500,
		Message: err.Error(),
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		httpError.Code = 404
	} else {
		switch e := err.(type) {
		case *fiber.Error:
			httpError.Code = e.Code
		case *ErrorDetail:
			httpError.Code = 400
			httpError.Detail = e
		case fiber.MultiError:
			httpError.Code = 400
			httpError.Message = ""
			for _, err = range e {
				httpError.Message += err.Error() + "\n"
			}
		}
	}

	return ctx.Status(httpError.Code).JSON(&httpError)
}
