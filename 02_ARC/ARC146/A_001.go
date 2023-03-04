package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	tmp := []int{1, 2, 3}

	ans := ""
	for i := 0; i < 7; i++ {
		newtmp := ""
		for _, j := range tmp {
			newtmp += strconv.Itoa(a[n-j])
		}
		ans = max(ans, newtmp)
		nextPermutation(sort.IntSlice(tmp))
	}
	fmt.Println(ans)
}

func max(a, b string) string {
	if a > b {
		return a
	}
	return b
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
