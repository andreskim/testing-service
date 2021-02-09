package main

import (
	"os"
	"testing"
)

func TestConfigure(t *testing.T) {
	e := configure()

	if e == nil {
		t.Error("Failed to create echo server")
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
