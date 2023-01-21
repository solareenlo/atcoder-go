package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)

	tree := redblacktree.NewWithIntComparator()
	for i := 0; i < q; i++ {
		var a int
		fmt.Fscan(in, &a)
		switch a {
		case 1:
			var x int
			fmt.Fscan(in, &x)
			val, found := tree.Get(x)
			if found {
				tree.Put(x, val.(int)+1)
			} else {
				tree.Put(x, 1)
			}
		case 2:
			var x, c int
			fmt.Fscan(in, &x, &c)
			val, found := tree.Get(x)
			if found {
				if val.(int) > c {
					tree.Put(x, val.(int)-c)
				} else {
					tree.Remove(x)
				}
			}
		case 3:
			fmt.Fprintln(out, tree.Right().Key.(int)-tree.Left().Key.(int))
		}
	}
}
