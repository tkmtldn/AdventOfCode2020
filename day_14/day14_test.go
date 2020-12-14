package main

import (
	"path/filepath"
	"testing"
)

func TestDay14(t *testing.T){
	path := filepath.Join(".", "test.txt")
	inp := ReadData(path)
	path2 := filepath.Join(".", "test2.txt")
	inp2 := ReadData(path2)

	res1 := DockingData(inp)
	exp1 := 165
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 14 Part 1: got %v want %v", res1, exp1)
	}

	res2 := DockingData2(inp2)
	exp2 := 208
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 14 Part 2: got %v want %v", res2, exp2)
	}
}