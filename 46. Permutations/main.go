package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}

	var res [][]int

	for i := range nums {

		for _, p := range permute(append(append([]int{}, nums[:i]...), nums[i+1:]...)) {
			res = append(res, append([]int{nums[i]}, p...))
		}
	}

	return res
}
