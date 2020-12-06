package main

import (
	"path/filepath"
	"testing"
)

func TestDay06(t *testing.T){
	path := filepath.Join(".", "day06testinput.txt")
	n :=ReadData(path)

	result_1 := CountElements(n, Count1)
	expected_1 := 11
	if result_1 != expected_1 {
		t.Errorf("unexpected result on Day 6 Part 1: got %v want %v", result_1, expected_1)
	}

	result_2 := CountElements(n, Count2)
	expected_2 := 6
	if result_2 != expected_2 {
		t.Errorf("unexpected result on Day 6 Part 2: got %v want %v", result_2, expected_2)
	}
}