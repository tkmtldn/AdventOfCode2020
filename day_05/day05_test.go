package main

import "testing"

func TestDay05(t *testing.T) {
	tests := []struct {
		input  string
		want int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}
	for _, tt := range tests {
		t.Run("SeatIDSearch", func(t *testing.T){
			got := SeatIDSearch(tt.input)
			if got != tt.want{
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}
}