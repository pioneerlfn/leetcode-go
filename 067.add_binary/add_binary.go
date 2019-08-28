package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	var sum, carry, remaider int
	var res bytes.Buffer
	for i >= 0 || j >= 0 {
		sum = carry
		if i >= 0 {
			sum += int(a[i]-'0')
			fmt.Println("sum:",sum)

		}
		if j >= 0 {
			sum += int(b[j]-'0')
			fmt.Println("sum:",sum)
		}
		carry, remaider = sum / 2, sum % 2
		res.WriteString(strconv.Itoa(remaider))
		i--
		j--
	}

	if carry != 0 {
		res.WriteString("1")
	}

	return reverseString(res.String())
}

func reverseString(s string) string {
	length := len(s)
	buf := []byte(s)
	for i := 0; i < length/2; i++ {
		buf[i], buf[length-1-i] = buf[length-1-i], buf[i]
	}

	return string(buf)
}

func main() {
	a := "0"
	b := "0"
	res := addBinary(a, b)
	fmt.Println(res)

	r := strings.NewReader("foo")


}