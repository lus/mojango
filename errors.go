package mojango

import (
	"errors"
	"github.com/valyala/fasthttp"
	"strconv"
)

// Define possible known errors
var ErrTooManyRequests = errors.New("too many Mojang API requests")
var ErrNoContent = errors.New("no Mojang API result")

// errorFromCode returns a corresponding error to the given status code
func errorFromCode(statusCode int) error {
	switch statusCode {
	case fasthttp.StatusTooManyRequests:
		return ErrTooManyRequests
	case fasthttp.StatusNoContent:
		return ErrNoContent
	default:
		return errors.New(strconv.Itoa(statusCode))
	}
}
