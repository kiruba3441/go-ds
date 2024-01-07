// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	input := []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}
	//fmt.Println(binary_search_recursion(0, len(input)-1, 1, input))
	//fmt.Println(binary_search_iterative(input, 1))
	fmt.Println(binary_search_rotated(0, len(input)-1, -5, input))

}

func binary_search_recursion(start, end, el int, input []int) int {
	if start > end || end >= len(input) || start < 0 {
		return -1
	}
	mid := start + (end-start)/2
	if el == input[mid] {
		return mid
	}

	if el > input[mid] {
		return binary_search_recursion(mid+1, end, el, input)
	} else {
		return binary_search_recursion(start, mid-1, el, input)
	}
}

func binary_search_iterative(input []int, el int) int {
	start := 0
	end := len(input) - 1
	for start <= end {
		mid := start + (end-start)/2
		if input[mid] == el {
			return mid
		}
		if el > input[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// 56789101234
// find mid point
// if matches all good
// no match
// 1. look left -> perfectly sorted & number falls in between -> search there
// 2. look right -> perfectly sorted & number falls in between -> search there
// 3. look for unsorted array -> search there.

func binary_search_rotated(start, end, el int, input []int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if input[mid] == el {
		return mid
	}
	left_array_st := start
	left_array_end := mid - 1
	if input[left_array_st] <= input[left_array_end] && el >= input[left_array_st] && el <= input[left_array_end] {
		return binary_search_rotated(left_array_st, left_array_end, el, input)
	}
	right_array_st := mid + 1
	right_array_end := end
	if input[right_array_st] <= input[right_array_end] && el >= input[right_array_st] && el <= input[right_array_end] {
		return binary_search_rotated(right_array_st, right_array_end, el, input)
	}
	if input[left_array_st] > input[left_array_end] {
		return binary_search_rotated(left_array_st, left_array_end, el, input)
	}
	if input[right_array_st] > input[right_array_end] {
		return binary_search_rotated(right_array_st, right_array_end, el, input)
	}
	return -1
}
