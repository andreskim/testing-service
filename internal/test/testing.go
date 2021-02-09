package internal_test

import (
	"fmt"
	"testing"
)

// TestError Panic on error
func TestError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		return
	}
	t.Fatalf("Error: %s", err)
}

// TestInfo Log error
func TestInfo(t *testing.T, format string, args ...interface{}) {
	t.Helper()

	t.Logf("%s", fmt.Sprintf(format, args...))
}
