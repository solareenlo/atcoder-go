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
	t := reverseString(s)
	n := len(s)

	f := ZAlgorithm(s)
	b := ZAlgorithm(t)
	b = append(b, 1<<60)
	res := 0
	for i := 0; i < n; i++ {
		r := n - i
		a := min(f[i], r-1)
		c := min(b[r], r-1)
		if r > a+c || i <= r {
			continue
		}
		res += a + c - r + 1
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func ZAlgorithm(s string) []int {
	n := len(s)
	z := make([]int, n)
	z[0] = n
	for i, j := 1, 0; i < n; {
		for i+j < n && s[j:j+1] == s[i+j:i+j+1] {
			j++
		}
		z[i] = j
		if j == 0 {
			i++
			continue
		}
		k := 1
		for ; i+k < n && k+z[k] < j; k++ {
			z[i+k] = z[k]
		}
		i, j = i+k, j-k
	}
	return z
}
