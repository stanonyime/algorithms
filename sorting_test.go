package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{[]int{5, 3, 4, 1, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{40, 11, 22, 31, 52}, []int{11, 22, 31, 40, 52}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf(" with arr %v", tt.arr)
		t.Run(testname, func(t *testing.T) {
			ans := InsertionSortInt(tt.arr)

			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}
