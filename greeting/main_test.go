package main

import (
	"testing"
)

func TestDouble(t *testing.T) {
	if Double(7) != 14 {
		t.Error("Expected 7 * 2 to equal 14")
	}
}
