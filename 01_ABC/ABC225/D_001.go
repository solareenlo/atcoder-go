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

	var n, q int
	fmt.Fscan(in, &n, &q)

	nil := -1
	front := make([]int, n+1)
	back := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		front[i] = nil
		back[i] = nil
	}

	for i := 0; i < q; i++ {
		var op, x int
		fmt.Fscan(in, &op, &x)
		var y int
		switch op {
		case 1:
			fmt.Fscan(in, &y)
			back[x] = y
			front[y] = x
		case 2:
			fmt.Fscan(in, &y)
			back[x] = nil
			front[y] = nil
		default:
			for front[x] != nil {
				x = front[x]
			}
			res := make([]int, 0)
			for x != nil {
				res = append(res, x)
				x = back[x]
			}
			s := fmt.Sprint(res)
			fmt.Fprintln(out, len(res), s[1:len(s)-1])
		}
	}
}
