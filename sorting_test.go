package algorithms

import (
	"fmt"
	"testing"
)

// TestMaxArraySumKadane
func TestInsertionSort(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{5, 3, 4, 1, 2}, 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf(" with arr %v", tt.arr)
		t.Run(testname, func(t *testing.T) {
			ans := InsertionSortInt(tt.arr)

			if ans[0] != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}
