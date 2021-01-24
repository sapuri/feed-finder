package feedfinder_test

import (
	"testing"

	"github.com/sapuri/feed-finder/feedfinder"
)

func TestNew(t *testing.T) {
	t.Parallel()

	client := feedfinder.New()
	if client == nil {
		t.Errorf("client is nil")
	}
}
