package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var l, r int
	var s string
	fmt.Fscan(in, &l, &r, &s)

	tmp := reverseString(s[l-1 : r])
	fmt.Println(s[:l-1] + tmp + s[r:])
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
