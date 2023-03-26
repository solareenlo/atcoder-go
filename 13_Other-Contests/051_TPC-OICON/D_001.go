package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	vec := make([]int, 0)
	for i := 2; i <= 10000; i++ {
		OK := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				OK = false
			}
		}
		if OK == true {
			vec = append(vec, i*i*i)
		}
	}

	var n int
	fmt.Fscan(in, &n)
	v := make([]int, 100009)
	M := make(map[int]int)
	maxn := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &v[i])
		for j := 0; j < len(vec); j++ {
			for v[i]%vec[j] == 0 {
				v[i] /= vec[j]
			}
		}
		M[v[i]]++
		maxn = max(maxn, M[v[i]])
	}
	fmt.Println(maxn)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
