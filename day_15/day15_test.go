package main

import (
	"testing"
)

func TestDay15(t *testing.T){
	inp := "0,3,6"

	res1 := RambunctiousRecitation(inp, 2020)
	exp1 := "436"
	if res1 != exp1 {
		t.Errorf("unexpected result on Day 15 Part 1: got %v want %v", res1, exp1)
	}

	res2 := RambunctiousRecitation(inp, 30000000)
	exp2 := "175594"
	if res2 != exp2 {
		t.Errorf("unexpected result on Day 15 Part 2: got %v want %v", res2, exp2)
	}
}
