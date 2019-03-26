package main

import (
	"fmt"
	"sort"
)

func FindSumTwoNums(arr []int, target int) [][2]int {
	ret := make([][2]int, 0, len(arr)/2)
	small, big := 0, len(arr)-1
	sum := 0
	for small < big {
		sum = arr[small] + arr[big]
		if sum == target {
			ret = append(ret, [2]int{small, big})
			small++
			big--
		} else if sum < target {
			small++
		} else {
			big--
		}
	}
	return ret
}

func main() {
	arr := []int{23, 21, 23, 78, 12, 23, 1, 34, 12, 45}

	indexPreMap := make(map[int][]int)
	for i, v := range arr {
		if _, ok := indexPreMap[v]; !ok {
			indexPreMap[v] = []int{i}
		} else {
			indexPreMap[v] = append(indexPreMap[v], i)
		}
	}

	keys := make([]int, 0, len(indexPreMap))
	for k, _ := range indexPreMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Printf("Origin array: %v\n", arr)
	sort.Ints(arr)
	fmt.Printf("Sorted array: %v\n", arr)

	sum := 46
	result := FindSumTwoNums(arr, sum)
	if len(result) > 0 {
		for i, two := range result {
			if len(indexPreMap[arr[two[0]]]) > 0 && len(indexPreMap[arr[two[1]]]) > 0 {
				first := indexPreMap[arr[two[0]]][0]
				if len(indexPreMap[arr[two[0]]]) > 1 {
					indexPreMap[arr[two[0]]] = indexPreMap[arr[two[0]]][1:]
				}
				second := indexPreMap[arr[two[1]]][0]
				if len(indexPreMap[arr[two[1]]]) > 1 {
					indexPreMap[arr[two[1]]] = indexPreMap[arr[two[1]]][1:]
				}
				fmt.Printf("%d couples: origin index: [%d,%d],sorted index: [%d,%d], values:[%d,%d]\n", i, first, second, two[0], two[1], arr[two[0]], arr[two[1]])
			}
		}
	} else {
		fmt.Printf("not found two nums that sum is %d\n", sum)
	}
}
