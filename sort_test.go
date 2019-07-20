package sort_test

import (
	"sort"
	"testing"

	csort "github.com/rgalus/quicksort-go"
	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{8, 4, 2, 9, 0, 7, 5, 6, 1, 3}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sortedArr := csort.QuickSort(arr)

	assert.Equal(t, expected, sortedArr)
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{8, 4, 2, 9, 0, 7, 5, 6, 1, 3}
		csort.QuickSort(arr)
	}
}

func BenchmarkBuilInSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{8, 4, 2, 9, 0, 7, 5, 6, 1, 3}
		sort.Ints(arr)
	}
}
