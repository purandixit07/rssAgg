package main

import "testing"

func TestNumber(t *testing.T) {
	result := unitTestFunc(3)
	if result != "correct" {
		t.Errorf("Result was incorrect, got %s, want: %s", result, "correct")
	}
}
