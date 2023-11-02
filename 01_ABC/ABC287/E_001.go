package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 500010

var ch [N][26]int
var cnt int
var prt, dep, bj, num [N]int

func add(s string, az int) {
	x := 0
	for i := 0; i < len(s); i++ {
		j := int(s[i] - 'a')
		if ch[x][j] == 0 {
			cnt++
			ch[x][j] = cnt
			prt[ch[x][j]] = x
			dep[ch[x][j]] = i + 1
		}
		x = ch[x][j]
		bj[x]++
	}
	num[az] = x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		add(s, i)
	}
	for i := 1; i <= n; i++ {
		x := num[i]
		for bj[x] == 1 {
			x = prt[x]
		}
		fmt.Fprintln(out, dep[x])
	}
}
