package main

import (
	"fmt"
	"strings"

	"github.com/nicolito128/arrays/each"
)

func main() {
	nums := []int{3, 4, 1, 1, 9, 8, 7, 7, 6, 5, 5}
	nums2 := each.Int(nums, mult)
	fmt.Println(nums)
	fmt.Println(nums2)

	fmt.Printf("/*******/\n")

	str := []string{"wallace", "dog", "carl", "earth"}
	fmt.Println(each.String(str, upper))
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
