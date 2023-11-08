package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200200

	var ans, bns [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	var s string
	fmt.Fscan(in, &s)
	str := strings.Split(s, "")
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ans[i])
	}
	for i := len(str) - 1; i >= 0; i-- {
		if bns[ans[i]] != 0 {
			str[i], str[bns[ans[i]]] = str[bns[ans[i]]], str[i]
		}
		bns[ans[i]] = i
	}
	fmt.Println(strings.Join(str, ""))
}
