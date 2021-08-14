package main

import (
	"fmt"
)

func QuickSort(left int, right int, arr *[9]int) {
	l := left
	r := right

	pivot := arr[(left+right)/2]

	for l < r {
		for arr[l] < pivot {
			l++
		}

		for arr[r] > pivot {
			r--
		}

		if l >= r {
			break
		}

		arr[l], arr[r] = arr[r], arr[l]

		// 优化
		if arr[l] == pivot {
			r--
		}

		if arr[r] == pivot {
			l++
		}
	}

	if l == r {
		l++
		r--
	}

	if left < r {
		QuickSort(left, r, arr)
	}

	if right > l {
		QuickSort(l, right, arr)
	}
}
func main() {
	arr := [9]int{3, 6, 7, 2, 8, 1, 9, 4, 5}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println()
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
}
