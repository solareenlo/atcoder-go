package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200200

var (
	a  = make([]int, N)
	ch = make([][]int, N)
	rt int
	ln int
	s  bool
)

func dfs(t int) int {
	ret := 0
	ap := make([]bool, 20)
	for _, u := range ch[t] {
		if u != rt {
			ap[dfs(u)] = true
		}
	}
	if t == a[rt] && s {
		ap[ln] = true
	}
	for ; ap[ret]; ret++ {

	}
	if t == a[rt] {
		ln = ret
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		ch[a[i]] = append(ch[a[i]], i)
	}

	var t int
	used := make([]bool, N)
	for ; !used[t]; t = a[t] {
		used[t] = true
	}

	rt = t
	rn := dfs(rt)
	if rn == ln {
		s = true
		if rn != dfs(rt) {
			fmt.Print("IM")
		}
	}
	fmt.Println("POSSIBLE")
}
