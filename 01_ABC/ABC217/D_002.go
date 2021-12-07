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

	var l, q int
	fmt.Fscan(in, &l, &q)

	tree := rbt.NewWithIntComparator()
	tree.Put(0, 0)
	tree.Put(l, l)

	for i := 0; i < q; i++ {
		var c, x int
		fmt.Fscan(in, &c, &x)
		switch c {
		case 1:
			tree.Put(x, x)
		case 2:
			l, _ := tree.Floor(x)
			r, _ := tree.Ceiling(x)
			fmt.Fprintln(out, r.Value.(int)-l.Value.(int))
		}
	}
}
