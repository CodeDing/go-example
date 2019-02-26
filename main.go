package main

import (
	"fmt"
	"time"
)

/*
   Thread 1 output: 1 2 3
   Thread 2 output: a b c

   output: 1a2b3c
*/

func main() {

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	done1 := make(chan bool, 1)
	done2 := make(chan bool, 1)
	arr1 := []rune{'1', '2', '3'}
	arr2 := []rune{'a', 'b', 'c'}

	ch1 <- 1
	go func() {
		for {
			<-ch1
			fmt.Printf("%c", arr1[0])
			if len(arr1) == 1 {
				done1 <- true
				close(ch2)
				break
			}
			if len(arr1) > 1 {
				arr1 = arr1[1:]
			}
			ch2 <- 1
		}
	}()

	go func() {
		for {
			<-ch2
			fmt.Printf("%c", arr2[0])
			if len(arr2) == 1 {
				done2 <- true
				close(ch1)
				break
			}
			if len(arr2) > 1 {
				arr2 = arr2[1:]
			}
			ch1 <- 1
		}
	}()

	for {
		select {
		case <-done1:
			fmt.Println("\nThread 1 done")
		case <-done2:
			fmt.Println("\nThread 2 done")
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			break
		}
	}
}
