/*
  select机制
  1)select+case是用于阻塞监听goroutine的,如果没有case,就单单一个select{},则监听当前程序中的goroutine(注意:需要有真实的goroutine在跑,否则select{}会panic)
  2)select里有多个可执行case,则随机执行一个
  3)select配合for循环监听channel有没有故事发生。在这个场景下,break只是退出当前select而不是for,退出for需要用break TIP/goto的方式
  4)无缓冲的通道,则传值后立马close,则会在close之前阻塞,有缓冲的通道则即使close了也会继续让接收后面的值
  5)同一个通道多个goroutine进行关闭, 可用recover panic的方式来判断通道的关闭问题

  源码参考:
     src/runtime/select.go
	 顺序遍历case来寻找可执行的case
*/
package main

import (
	"fmt"
	"time"
)

/*
*  Example 1
 */
//func main() {
//	t1 := time.Tick(time.Second)
//	t2 := time.Tick(time.Second)
//	var count int
//	for {
//		select {
//		case <-t1:
//			fmt.Println("第一个case")
//			count++
//			fmt.Println("count==>", count)
//		case <-t2:
//			fmt.Println("第二个case")
//			count++
//			fmt.Println("count==>", count)
//
//		}
//	}
//}

func main() {

	var count int
	for {
		select {
		case <-time.Tick(time.Microsecond * 499):
			fmt.Println("第一个case")
			count++
			fmt.Println("count->", count)
		case <-time.Tick(time.Microsecond * 500):
			fmt.Println("第二个case")
			count++
			fmt.Println("count->", count)
		}
	}
}
