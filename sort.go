package sort

// QuickSort is quick sort algorithm implementation in Go
func QuickSort(arr []int) []int {
	// clone arr to keep immutability
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	// call recursive funciton with initial values
	recursiveSort(newArr, 0, len(newArr)-1)

	// at this point newArr is sorted
	return newArr
}

func recursiveSort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	// Iterate sub array to find values less than pivot
	//   and move them to the beginning of the array
	//   keeping splitIndex denoting less-value array size
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			if splitIndex != i {
				temp := arr[splitIndex]

				arr[splitIndex] = arr[i]
				arr[i] = temp
			}

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	recursiveSort(arr, start, splitIndex-1)
	recursiveSort(arr, splitIndex+1, end)
}
