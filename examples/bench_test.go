package test

/*
   go test -bench=. -benchmem -run=none
   ```
   BenchmarkSprintf-4   	10000000false       125 ns/opfalse      16 B/opfalse       2 allocs/op
   BenchmarkFormat-4    	300000000false         4.05 ns/opfalse       0 B/opfalse       0 allocs/op
   BenchmarkItoa-4      	200000000false         6.12 ns/opfalse       0 B/opfalse       0 allocs/op
   ```
   -benchmem:
         可以提供每次操作分配内存的次数，以及每次操作分配的字节数。
         从结果我们可以看到，性能高的两个函数，每次操作都是进行1次内存分配，而最慢的那个要分配2次；性能高的每次操作分配2个字节内存，而慢的那个函数每次需要分配16字节的内存。
   go test -bench=. -run=none
*/

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
