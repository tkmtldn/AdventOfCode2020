package main

import (
	"path/filepath"
	"testing"
)

func TestDay08(t *testing.T) {
	path := filepath.Join(".", "test.txt")
	n := ReadData(path)

	res1 := Accumulator(n)
	exp1 := 5
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 8 Part 1: got %v want %v", res1, exp1)
	}

	res2 := Accumulator2(n)
	exp2 := 8
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 8 Part 2: got %v want %v", res2, exp2)
	}
}