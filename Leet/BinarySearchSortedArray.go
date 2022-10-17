package Leet

import "fmt"

func FindIndexSortedArray() {
	array := []int{1, 2, 3, 5, 6, 7}

	fmt.Println("Hello to Go FindIndexSortedArray")
	fmt.Println("the sorted array is ", array)
	fmt.Println("Enter the value to search: ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Wrong input")
		fmt.Println(err)
		return
	}

	fmt.Println("looking for", input)
	index := binarySearchSortedArray(array, input)

	if -1 == index {
		fmt.Println("not Found")
	} else {
		fmt.Println("index is", index)
	}

}

func binarySearchSortedArray(nums []int, target int) int {
	left := 0
	right := len(nums)
	maxIndex := right - 1
	//perform binary search
	for left <= right {
		mid := (left + right) / 2
		switch {
		case mid > maxIndex:
			return -1
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}
