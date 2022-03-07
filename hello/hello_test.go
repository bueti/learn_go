package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ben")
	want := "Hello, Ben"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
