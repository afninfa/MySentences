package main

import "testing"

func ShouldError(t *testing.T, result error) {
	if result == nil {
		t.Errorf("expected non-nil value, got nil")
	}
}

func ShouldNotError(t *testing.T, result error) {
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}
