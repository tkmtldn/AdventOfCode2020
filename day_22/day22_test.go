package main

import (
	"path/filepath"
	"testing"
)

func TestDay22(t *testing.T) {
	path := filepath.Join(".", "test.txt")
	inp1, inp2 := ReadData(path)

	res1 := CrabCombat(inp1, inp2)
	exp1 := 306
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 22 Part 1: got %v want %v", res1, exp1)
	}

	res2 := CrabCombat2(inp1, inp2)
	exp2 := 291
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 22 Part 2: got %v want %v", res2, exp2)
	}
}