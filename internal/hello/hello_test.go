package goplay

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	fmt.Println(want)
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
