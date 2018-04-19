package main

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestResumeFail(t *testing.T) {
	if err := AssertThat(8080, Is(8080)); err != nil {
		t.Fatal(err)
	}
}
