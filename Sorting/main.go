package main

import (
	// "Sorting/SortingAlgos"
	sortingquestions "Sorting/SortingQuestions"
	"fmt"
)

func main() {
	arr := []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1}

	// sorted_arr := SortingAlgos.CountSort(arr)
	// SortingAlgos.QuickSort(arr)
	// fmt.Println("Sorted array:", arr)
	fmt.Println(sortingquestions.Partition01(arr, len(arr)))
}
