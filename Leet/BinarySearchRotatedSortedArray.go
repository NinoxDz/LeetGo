package Leet

import "fmt"

func FindIndexRotatedSortedArray() {
	array := []int{10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 5, 6, 7, 8, 9}

	fmt.Println("Hello to Go FindIndexRotatedSortedArray")
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
	fmt.Println("index", index)

	if -1 == index {
		fmt.Println("not Found")
	} else {
		fmt.Println("index is", index)
	}
}

func binarySearchRotatedSortedArray(nums []int, target int) int {
	pivot := getPivotLocation(nums)
	left, right := 0, len(nums)-1
	maxIndex := right - 1

	switch {
	case pivot == -1: //array is not rotated, search the entire array
	case target > nums[right]: //if target is in the left array
		right = pivot
	default: //if target in the right array
		left = pivot
	}

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

// get pivot location
func getPivotLocation(nums []int) int {

	left, right := 0, len(nums)-1
	//non-rotated array
	if nums[left] < nums[right] {
		return -1
	}
	for left <= right {
		mid := (left + right) / 2

		switch {
		case mid < len(nums)-1 && nums[mid] > nums[mid+1]:
			return mid + 1
		case mid > 0 && nums[mid] < nums[mid-1]:
			return mid
		case nums[left] < nums[mid]:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}
