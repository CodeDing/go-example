package main

import (
	"fmt"
	"strconv"
)

func checkValidBinaryValue(vs ...string) bool {
	for _, v := range vs {
		for _, b := range v {
			if b != '0' && b != '1' {
				return false
			}
		}
	}
	return true
}

func addBinary(a, b string) (string, error) {
	if len(a) < 1 || len(b) < 1 {
		return "", fmt.Errorf("a and b must not be empty!")
	}

	if !checkValidBinaryValue(a, b) {
		return "", fmt.Errorf("invalid a or b, elem in a or b must be '0' or '1'")
	}
	result := ""
	alen := len(a)
	blen := len(b)
	var abit, bbit, carrier int
	for alen > 0 || blen > 0 {
		if alen > 0 {
			abit = int(a[alen-1] - '0')
		} else {
			abit = 0
		}
		if blen > 0 {
			bbit = int(b[blen-1] - '0')
		} else {
			bbit = 0
		}

		bit := strconv.Itoa((abit + bbit + carrier) % 2)
		result = bit + result
		if abit+bbit+carrier > 1 {
			carrier = 1
		} else {
			carrier = 0
		}
		//fmt.Printf("alen=%d, blen=%d, abit=%d, bbit=%d, carrier=%d, result=%s\n", alen, blen, abit, bbit, carrier, result)
		alen--
		blen--
	}
	if carrier == 1 {
		result = "1" + result
	}
	return result, nil

}

func main() {
	as := [...]string{"1111", "11001"}
	bs := [...]string{"110", "1101"}
	for _, a := range as {
		for _, b := range bs {
			ret, _ := addBinary(a, b)
			fmt.Printf("Binary(add): %s + %s = %s\n", a, b, ret)
		}
	}

}
