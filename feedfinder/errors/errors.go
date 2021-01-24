package errors

import (
	"fmt"
)

type HTTPError error

func NewHTTPError(statusCode int) HTTPError {
	return fmt.Errorf("HTTP error: returned %d", statusCode)
}
