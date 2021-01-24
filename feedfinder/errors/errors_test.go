package errors_test

import (
	"net/http"
	"testing"

	"github.com/sapuri/feed-finder/feedfinder/errors"
)

func TestNewHTTPError(t *testing.T) {
	t.Parallel()

	want := "HTTP error: returned 404"
	if got := errors.NewHTTPError(http.StatusNotFound).Error(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
