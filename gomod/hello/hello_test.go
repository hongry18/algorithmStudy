package hello

import (
    "testing"
    "myLib/sample"
)

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

func TestSample(t *testing.T) {
    want := "sample"
    if got := sample.Sample(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
