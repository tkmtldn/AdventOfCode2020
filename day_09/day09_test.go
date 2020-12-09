package main

import (
	"path/filepath"
	"testing"
)

func TestDay09(t *testing.T) {
	path := filepath.Join(".", "test.txt")
	n := ReadData(path)

	res1 := PreambleC(n, 5)
	exp1 := 127
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 9 Part 1: got %v want %v", res1, exp1)
	}

	res2 := PreambleWeakness(n, res1)
	exp2 := 62
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 9 Part 2: got %v want %v", res2, exp2)
	}
}
