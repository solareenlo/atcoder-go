package main

import (
	"bufio"
	"fmt"
	"os"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var Q int
	fmt.Fscan(in, &Q)

	tree := rbt.NewWithIntComparator()
	for i := 0; i < Q; i++ {
		var t, x int
		fmt.Fscan(in, &t, &x)
		switch t {
		case 1:
			tree.Put(x, x)
		case 2:
			l, ok1 := tree.Floor(x)
			r, ok2 := tree.Ceiling(x)
			if !ok1 && !ok2 {
				fmt.Fprintln(out, -1)
			} else if !ok1 {
				fmt.Fprintln(out, r.Value.(int)-x)
			} else if !ok2 {
				fmt.Fprintln(out, x-l.Value.(int))
			} else {
				fmt.Fprintln(out, min(r.Value.(int)-x, x-l.Value.(int)))
			}

		}

	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
