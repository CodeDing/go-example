package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	var content string
	for _, v := range nums {
		content += fmt.Sprintf("%d, ", v)
	}
	content = content[:len(content)-2]
	for i := 0; i <= len(nums)-3; i++ {
		j := i + 1
		k := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for ; j < k && nums[j] == nums[j-1]; j++ {

				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1, 1, -2, -3, 5, 0, 0, 0, 7, 8}
	var content string
	for _, v := range nums {
		content += fmt.Sprintf("%d, ", v)
	}
	content = content[:len(content)-2]
	fmt.Printf("origin number: \n{%s}\n", content)
	slices := threeSum(nums)
	fmt.Println("result:")
	for _, arr := range slices {
		content := content[:0]
		for _, v := range arr {
			content += fmt.Sprintf("%d, ", v)
		}
		content = content[:len(content)-2]
		fmt.Printf("{%s}\n", content)
	}
}
