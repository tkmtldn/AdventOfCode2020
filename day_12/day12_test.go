package main

import (
	"path/filepath"
	"testing"
)

func TestDay12(t *testing.T){
	path := filepath.Join(".", "test.txt")
	inp := ReadData(path)

	res1 := CalcDist(inp)
	exp1 := 25
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 12 Part 1: got %v want %v", res1, exp1)
	}

	res2 := CalcDistToo(inp)
	exp2 := 286
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 12 Part 2: got %v want %v", res2, exp2)
	}
}