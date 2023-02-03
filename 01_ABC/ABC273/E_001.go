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

	const maxn = 500050

	var q int
	fmt.Fscan(in, &q)

	var val, fa [maxn]int
	val[0] = -1
	tot, cur := 0, 0
	sv := make(map[int]int)
	for q > 0 {
		var op string
		fmt.Fscan(in, &op)
		switch op {
		case "ADD":
			var x int
			fmt.Fscan(in, &x)
			tot++
			fa[tot] = cur
			cur = tot
			val[cur] = x
		case "DELETE":
			cur = fa[cur]
		case "SAVE":
			var y int
			fmt.Fscan(in, &y)
			sv[y] = cur
		case "LOAD":
			var y int
			fmt.Fscan(in, &y)
			cur = sv[y]
		}
		fmt.Fprintf(out, "%d ", val[cur])
		q--
	}
}
