package main

import (
	"path/filepath"
	"testing"
)

func TestDay06(t *testing.T){
	path := filepath.Join(".", "test.txt")
	n :=ReadData(path)

	res1 := CountElements(n, Count1)
	exp1 := 11
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 6 Part 1: got %v want %v", res1, exp1)
	}

	res2 := CountElements(n, Count2)
	exp2 := 6
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 6 Part 2: got %v want %v", res2, exp2)
	}
}