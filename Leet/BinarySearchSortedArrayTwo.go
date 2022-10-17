package Leet

import "fmt"

func FindIndexSortedArrayTwo() {
	array := []int{1, 2, 2, 2, 3, 3, 3, 3, 3, 6, 7, 7}

	fmt.Println("Hello to Go FindIndexSortedArrayTwo")
	fmt.Println("This Function will return the index of a value if it exist")
	fmt.Println("or where it should be if it doesn't exist")
	fmt.Println("to ether help you find the place to insert it or get the lenth of te target value -1")
	fmt.Println("the sorted array is ", array)
	fmt.Println("Enter the value to get it's length: ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Wrong input")
		fmt.Println(err)
		return
	}

	fmt.Println("looking for", input)
	indexFirst := binarySearchTwo(array, input)
	indexSecond := binarySearchTwo(array, input+1)

	fmt.Println("the length of", input, "is", indexSecond-indexFirst)

}

func binarySearchTwo(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
