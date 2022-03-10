package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	s := make([]int, n)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}
	t := make([]int, m)
	for i := range t {
		fmt.Fscan(in, &t[i])
	}

	S := n
	T := m
	sort.Ints(s)
	sort.Ints(t)

	const mod = 1_000_000_007
	ans := 1
	for i := n * m; i > 0; i-- {
		if S-1 >= 0 && T-1 >= 0 && s[S-1] == i && t[T-1] == i && S*T != 0 {
			S--
			T--
		} else if S-1 >= 0 && s[S-1] == i && S != 0 {
			ans *= m - T
			S--
		} else if T-1 >= 0 && t[T-1] == i && T != 0 {
			ans *= n - S
			T--
		} else {
			ans *= (n-S)*(m-T) - (n*m - i)
		}
		ans %= mod
	}
	fmt.Println(ans)
}
