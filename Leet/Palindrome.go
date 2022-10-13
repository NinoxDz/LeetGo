package Leet

import (
	"fmt"
	"strconv"
)

func Palindrome(number int) bool {
	fmt.Println("Checking the number if Palindrome", number)
	stringNumber := strconv.Itoa(number)

	chars := []rune(stringNumber)
	lenth := len(chars)
	i := 0
	j := lenth
	for i < lenth/2 {
		if chars[i] != chars[j-1] {
			fmt.Println("false the number is not Palindrome")

			return false
		}

		i++
		j--
	}
	fmt.Println("true the number is Palindrome")

	return true
}
