/*
	数组(非递减),找出给定元素出现的最后一个索引
	如:
	    [1,3,5,5,5,5,5,5,8,9]
		findIndex(arr,5) = 8
*/
package main

import "fmt"

func findIndexOfKey(arr []int, key int) int {

	var flag bool
	for i, v := range arr {
		if v == key {
			flag = true
		}
		if flag && v != key {
			return i - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 5, 5, 5, 5, 5, 8, 9}
	key := 5
	idx := findIndexOfKey(arr, key)
	fmt.Printf("Origin of array : %v\n", arr)
	fmt.Printf("Index of key %d is %d\n", key, idx)
}
