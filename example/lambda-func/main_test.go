package main
package main

import (
	"context"
	"testing"
)

func TestMyHandler(t *testing.T) {
	// Initialize mock input payload matching your handler signature
	input := MyEvent{Key: "test-value"}

	// Invoke the handler directly in pure Go
	response, err := HandleRequest(context.Background(), input)

	// Validate outcomes
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Message != "Success" {
		t.Errorf("Expected 'Success', got %s", response.Message)
	}
}
