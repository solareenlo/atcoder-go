package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	n := make([]int, 0)
	for i := 0; i < len(s); i++ {
		n = append(n, int(s[i]-'0'))
	}

	var k int
	fmt.Fscan(in, &k)
	for k > 0 {
		k--
		m := 0
		for i := 0; i < len(n); i++ {
			m = m*8 + n[i]
		}
		n = make([]int, 0)
		if m == 0 {
			n = append(n, 0)
		}
		for m > 0 {
			n = append(n, m%9)
			m /= 9
		}
		n = reverseOrderInt(n)
		for i := 0; i < len(n); i++ {
			if n[i] == 8 {
				n[i] = 5
			}
		}
	}
	for _, i := range n {
		fmt.Print(i)
	}
	fmt.Println()
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
