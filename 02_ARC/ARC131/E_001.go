package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	i := n
	if n < 5 || n%3 == 2 {
		fmt.Fprintln(out, "No")
		return
	}

	a := [3]int{}
	fmt.Fprintln(out, "Yes")
	j := 0
	i--
	for ; i >= 0; i-- {
		for k := 2; k >= 0; k-- {
			if a[k]+i <= n*(n-1)/6 {
				a[k] += i
				j = k
				break
			}
		}
		t := "BRW"
		fmt.Fprintln(out, strings.Repeat(string(t[j]), i))
	}
}
