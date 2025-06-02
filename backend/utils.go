package main

import "testing"

func Expect(e error) {
	if e != nil {
		panic(e)
	}
}

func Unwrap[T any](t T, e error) T {
	Expect(e)
	return t
}

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
