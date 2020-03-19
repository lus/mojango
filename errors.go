package mojango

import (
	"errors"
	"github.com/valyala/fasthttp"
	"strconv"
)

// Define possible known errors
var ErrTooManyRequests = errors.New("too many requests")

// Returns a corresponding error to the given status code
func errorFromCode(statusCode int) error {
	switch statusCode {
	case fasthttp.StatusTooManyRequests:
		return ErrTooManyRequests
	default:
		return errors.New(strconv.Itoa(statusCode))
	}
}
