package main

import "fmt"

func SelectionSort(arr *[9]int) {

	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		minIdx := i
		for j := i + 1; j < len(*arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
func main() {
	arr := [9]int{3, 6, 7, 2, 8, 1, 9, 4, 5}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	SelectionSort(&arr)
	fmt.Println()
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}

}
