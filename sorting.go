package algorithms

// InsertionSortInt returns int slice sorted in ascending order
func InsertionSortInt(arr []int) []int {
	var min int

	for i := 0; i < len(arr); i++ {
		min = i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}

	return arr

}
