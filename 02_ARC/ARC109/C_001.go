package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(x, y byte) byte {
	z := string(x) + string(y)
	if z == "RP" || z == "PS" || z == "SR" {
		return y
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	var t string
	fmt.Fscan(in, &n, &k, &t)

	for j := 0; j < k; j++ {
		s := t + t
		t = ""
		for i := 0; i < n; i++ {
			t += string(f(s[i*2], s[i*2+1]))
		}
	}
	fmt.Println(string(t[0]))
}
