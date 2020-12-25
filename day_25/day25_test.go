package main

import (
	"testing"
)

func TestDay25(t *testing.T) {

	res := ComboBreaker(5764801, 17807724)
	exp := 14897079
	if res != exp {
		t.Errorf("unexpected result on Day 25 Part 1: got %v want %v", res, exp)
	}
}
