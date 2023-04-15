package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var strs [16]string
	var m int
	fmt.Fscan(in, &strs[0], &m)
	for i := 1; i < 16; i++ {
		strs[i] = magic(strs[i-1])
	}
	t := 0
	for m > 0 {
		m--
		var x int
		var s string
		fmt.Fscan(in, &x, &s)
		for i := 0; i < 16 && s >= strs[t%16]; i++ {
			t++
		}
		if s >= strs[t%16] {
			fmt.Println("No")
			return
		}
		if t > x {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	return
}

func magic(a string) string {
	a = "00" + a
	res := make([]string, 0)
	for i := 2; i < len(a); i++ {
		if a[i] == a[i-2] {
			res = append(res, "0")
		} else {
			res = append(res, "1")
		}
	}
	return strings.Join(res, "")
}
