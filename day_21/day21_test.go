package main

import (
	"path/filepath"
	"testing"
)

func TestDay21(t *testing.T){
	path := filepath.Join(".", "test.txt")
	ingred, vocab := ReadData(path)
	res1, res2 := AllergenAssessment(ingred, vocab)
	exp1 := 5
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 21 Part 1: got %v want %v", res1, exp1)
	}
	exp2 := "mxmxvkd,sqjhc,fvjkl"
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 21 Part 2: got %v want %v", res2, exp2)
	}
}
