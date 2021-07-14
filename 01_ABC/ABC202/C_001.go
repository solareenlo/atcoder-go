package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	b := make([]int, n)
	for i := range b {
		fmt.Fscan(in, &b[i])
		b[i]--
	}
	c := make([]int, n)
	for i := range c {
		fmt.Fscan(in, &c[i])
		c[i]--
	}
	cnt := make([]int, n)
	for i := range cnt {
		cnt[b[c[i]]]++
	}
	res := 0
	for i := range cnt {
		res += cnt[a[i]]
	}
	fmt.Println(res)
}
