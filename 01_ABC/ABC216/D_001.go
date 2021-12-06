package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(stdin, &n, &m)

	a := make([]int, 2*n+1)
	in := make([]int, n+1)
	v := make([][]int, n+1)
	for i := 1; i < m+1; i++ {
		var k int
		fmt.Fscan(stdin, &k)
		for j := 0; j < k; j++ {
			fmt.Fscan(stdin, &a[j])
			if j != 0 {
				v[a[j-1]] = append(v[a[j-1]], a[j])
				in[a[j]]++
			}
		}
	}

	q := make([]int, 0)
	for i := 1; i < n+1; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		n--
		for _, i := range v[u] {
			in[i]--
			if in[i] == 0 {
				q = append(q, i)
			}
		}
	}

	if n != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
