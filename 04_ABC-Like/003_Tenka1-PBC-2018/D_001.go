package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	k := 0
	for k*(k-1)/2 < n {
		k++
	}

	if k*(k-1)/2 != n {
		fmt.Fprintln(out, "No")
		return
	}

	cnt := 0
	f := [1005][1005]int{}
	l := [1005]int{}
	fmt.Fprintln(out, "Yes")
	fmt.Fprintln(out, k)
	for i := 1; i < k; i++ {
		for j := k; j > i; j-- {
			cnt++
			l[i]++
			f[i][l[i]] = cnt
			l[j]++
			f[j][l[j]] = cnt
		}
	}

	for i := 1; i <= k; i++ {
		fmt.Fprint(out, l[i], " ")
		for j := 1; j < l[i]; j++ {
			fmt.Fprint(out, f[i][j], " ")
		}
		fmt.Fprintln(out, f[i][l[i]])
	}
}
