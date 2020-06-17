package algorithms

import (
	"fmt"
	"testing"
)

// TestMaxArraySumKadane
func TestMaxArraySumKadane(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{34, -50, 42, 14, -5, 86}, 137},
		{[]int{-5, -1, -8, -9}, 0},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf(" with arr %v", tt.arr)
		t.Run(testname, func(t *testing.T) {
			ans := MaxArraySumKadane(tt.arr)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
