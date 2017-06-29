package status

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsHandler(t *testing.T) {
	object := NewHandler("/")
	var expected *http.Handler
	err := AssertThat(object, Implements(expected))
	if err != nil {
		t.Fatal(err)
	}
}
