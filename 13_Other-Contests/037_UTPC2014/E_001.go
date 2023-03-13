package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	to  [10]int
	sum int
	max int
}

var v []node
var t node
var a string
var b int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	v = make([]node, 2)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		a = reverseString(a)
		fmt.Fprintln(out, dfs(1, 0))
	}
}

func dfs(n, k int) int {
	if k == len(a) {
		v[n].sum += b
		return v[n].max + v[n].sum
	}
	ta := a[k] - '0'
	if v[n].to[ta] == 0 {
		v[n].to[ta] = len(v)
		v = append(v, t)
	}
	tmp := dfs(v[n].to[ta], k+1)
	v[n].max = max(v[n].max, tmp)
	return v[n].sum + v[n].max
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
