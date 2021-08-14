package main

import "fmt"

func InsertionSort(arr *[9]int) {

	for i := 1; i < len(arr); i++ {

		insertVal := arr[i]
		insertPosition := i - 1

		for insertPosition >= 0 && arr[insertPosition] > insertVal {
			arr[insertPosition+1] = arr[insertPosition]
			insertPosition--
		}

		// 如果insertPosition没有移动, 即insertPosition + 1 == i, 则此元素不需要进行排序
		if insertPosition+1 != i {
			arr[insertPosition+1] = insertVal
		}

		fmt.Printf("After %d times, arr = %v\n", i, arr)
	}
}

func main() {
	arr := [9]int{3, 6, 7, 2, 8, 1, 9, 4, 5}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	InsertionSort(&arr)
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
}
