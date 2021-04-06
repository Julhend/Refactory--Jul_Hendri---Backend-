package main

import (
	"fmt"
)

func main() {
	fmt.Println(maskLeft("23dn3ir30fd2eddd"))
}

func maskLeft(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)-3; i++ {
		rs[i] = '*'
	}
	return string(rs)
}
