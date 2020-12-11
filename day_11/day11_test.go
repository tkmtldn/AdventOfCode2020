package main

import (
	"path/filepath"
	"testing"
)

func TestDay11(t *testing.T){
	path := filepath.Join(".", "test.txt")
	inp, x, y := ReadData(path)

	res1 := SeatingSystem(inp, x, y, Occupied)
	exp1 := 37
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 11 Part 1: got %v want %v", res1, exp1)
	}

	res2 := SeatingSystem(inp, x, y, OccupiedTwo)
	exp2 := 26
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 11 Part 2: got %v want %v", res2, exp2)
	}
}