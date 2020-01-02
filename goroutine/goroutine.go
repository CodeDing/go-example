package main

import "fmt"

//Thread 1 output: 1, 2, 3
//Thread 2 output: a, b, c
func main() {

	nums := []rune{'1', '2', '3'}
	chs := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	c := make(chan rune)
	go printChars(chs, c)
	var content string
	var num rune
	idx := 0
	for v := range c {
		if idx < len(nums) {
			num = nums[idx]
		} else {
			num = 0
		}
		content += fmt.Sprintf("%c%c", v, num)
		idx++
	}
	fmt.Printf("%s\n", content)
}

func printChars(chs []rune, c chan<- rune) {
	for i := 0; i < len(chs); i++ {
		c <- chs[i]
	}
	close(c)
}
