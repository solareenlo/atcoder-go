package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Scan(&n)
	aa := make([]int, n)
	bb := make([]int, n)
	cc := make([]int, n)
	for i := range aa {
		fmt.Fscan(in, &aa[i])
	}
	for i := range bb {
		fmt.Fscan(in, &bb[i])
	}
	for i := range cc {
		fmt.Fscan(in, &cc[i])
	}
	sort.Ints(aa)
	sort.Ints(bb)
	sort.Ints(cc)
	res := 0
	for i := 0; i < n; i++ {
		a := aa[i]
		bi := sort.SearchInts(bb, a+1)
		if len(bb) <= bi {
			break
		}
		b := bb[bi]
		bb = bb[bi+1:]
		ci := sort.SearchInts(cc, b+1)
		if len(cc) <= ci {
			break
		}
		cc = cc[ci+1:]
		res++
	}
	fmt.Println(res)
}
