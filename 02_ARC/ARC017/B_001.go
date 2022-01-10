package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	diff := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		if a[i+1]-a[i] > 0 {
			diff[i] = 1
		}
	}

	s := make([]int, n)
	for i := 0; i < n-1; i++ {
		s[i+1] = s[i] + diff[i]
	}

	cnt := 0
	for i := 0; i < n-k+1; i++ {
		if s[i+k-1]-s[i] == k-1 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
