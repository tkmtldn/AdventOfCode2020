package main

import (
	"path/filepath"
	"testing"
)

func TestDay10(t *testing.T) {
	path1 := filepath.Join(".", "test1.txt")
	path2 := filepath.Join(".", "test2.txt")
	inp1 := ReadData(path1)
	inp2 := ReadData(path2)

	res1 := AdapterDiff(inp1)
	exp1 := 35
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 10 Part 1: got %v want %v", res1, exp1)
	}

	res2 := AdapterComb(inp1)
	exp2 := 8
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 10 Part 2: got %v want %v", res2, exp2)
	}

	res3 := AdapterDiff(inp2)
	exp3 := 220
	if res3 != exp3 {
		t.Errorf("unexpected result on Day 10 Part 1: got %v want %v", res3, exp3)
	}

	res4 := AdapterComb(inp2)
	exp4 := 19208
	if res4 != exp4 {
		t.Errorf("unexpected result on Day 10 Part 2: got %v want %v", res4, exp4)
	}
}