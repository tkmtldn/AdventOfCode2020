package main

import (
	"testing"
)

func TestDay18(t *testing.T) {
	tests1 := []struct {
		input []string
		want  int
	}{
		{[]string{"1 + 2 * 3 + 4 * 5 + 6"}, 71},
		{[]string{"1 + (2 * 3) + (4 * (5 + 6))"}, 51},
		{[]string{"2 * 3 + (4 * 5)"}, 26},
		{[]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, 437},
		{[]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, 12240},
		{[]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, 13632},
	}
	for _, tt := range tests1 {
		t.Run("WithoutPriority", func(t *testing.T) {
			got := OperationOrder(tt.input, WithoutPriority)
			if got != tt.want {
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}

	tests2 := []struct {
		input []string
		want  int
	}{
		{[]string{"1 + 2 * 3 + 4 * 5 + 6"}, 231},
		{[]string{"1 + (2 * 3) + (4 * (5 + 6))"}, 51},
		{[]string{"2 * 3 + (4 * 5)"}, 46},
		{[]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, 1445},
		{[]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, 669060},
		{[]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, 23340},
	}
	for _, tt := range tests2 {
		t.Run("WithPriority", func(t *testing.T) {
			got := OperationOrder(tt.input, WithPriority)
			if got != tt.want {
				t.Errorf("Error when running with value %v, got %v, want %v.", tt.input, got, tt.want)
			}
		})
	}
}
