package helper

import (
	"errors"
	"testing"
)

func TestErrorPanic(t *testing.T) {
	// Test case 1: nil error should not panic
	ErrorPanic(nil) // This should not panic

	// Test case 2: non-nil error should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for non-nil error, but no panic occurred")
		}
	}()

	// This should trigger a panic
	ErrorPanic(errors.New("test error"))
}
