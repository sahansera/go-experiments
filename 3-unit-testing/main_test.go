package main

import "testing"

func TestAddOne(t *testing.T) {
	ans := AddOne(1)
	if ans != 2 {
		t.Errorf("Expected 2, got %d", ans)
	}
}
