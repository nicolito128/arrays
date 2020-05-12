package main

import (
	"fmt"
	"strings"

	"github.com/nicolito128/arrays/each"
	"github.com/nicolito128/arrays/filter"
)

func main() {
	nums := []int{3, 4, 1, 9, 9, 8, 7, 7, 6, 5, 5}
	nums2 := each.Int(nums, mult)
	fmt.Println(nums)
	fmt.Println(nums2)

	fmt.Printf("\n/*******/\n\n")

	str := []string{"wallace", "dog", "carl", "earth"}
	fmt.Println(each.String(str, upper))

	fmt.Printf("\n/*******/\n\n")

	evenNums := filter.Int(nums, filterEvenNumbers)
	divisibleByThree := filter.Int(nums, divByThree)
	fmt.Println(evenNums)
	fmt.Println(divisibleByThree)
}

func mult(el int, i int, arr []int) int {
	if el > 5 {
		return el * 3
	}

	return el * 2
}

func upper(el string, i int, arr []string) string {
	return strings.ToUpper(el)
}

func filterEvenNumbers(el int, i int, arr []int) bool {
	return (el % 2) == 0
}

func divByThree(el int, i int, arr []int) bool {
	return (el % 3) == 0
}
