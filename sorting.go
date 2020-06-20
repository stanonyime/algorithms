package algorithms

// InsertionSortInt returns int slice sorted in ascending order
func InsertionSortInt(arr []int) []int {
	var min int
	res := make([]int, len(arr), cap(arr))

	for i := 0; i < len(arr); i++ {
		min = i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		res[i] = arr[min]
		res[min] = arr[i]
	}

	return res

}
