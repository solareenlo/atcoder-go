package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tri [3003003][3]int

	var n, a int
	fmt.Fscan(in, &n, &a)
	cnt := 0
	for i := 0; i < n-1; i++ {
		var b int
		fmt.Fscan(in, &b)
		t := 0
		for j := 0; j < 30; j++ {
			k := (a >> (29 - j)) & 1
			if tri[t][k] == 0 {
				cnt++
				tri[t][k] = cnt
			}
			t = tri[t][k]
			tri[t][2]++
		}
		a ^= b
	}
	nw := 0
	for i := 0; i < n-1; i++ {
		nx, t := 0, 0
		for j := 0; j < 30; j++ {
			k := (nw >> (29 - j)) & 1
			if tri[tri[t][k]][2] == 0 {
				k ^= 1
			}
			nx |= k << (29 - j)
			t = tri[t][k]
			tri[t][2]--
		}
		fmt.Fprintf(out, "%d ", nx^nw)
		nw = nx
	}
	fmt.Fprintln(out, nw^a)
}
