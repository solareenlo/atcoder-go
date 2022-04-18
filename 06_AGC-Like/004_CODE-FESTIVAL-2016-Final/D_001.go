package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	f = [100100]int{}
	s = [100100]int{}
)

func cal(i, j int) int {
	if f[i] == f[j] {
		return f[i]
	}
	if f[i] > f[j] {
		i, j = j, i
	}
	return f[i] + min((f[j]-f[i])>>1, s[j])
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	g := make(map[int]int)
	for i := 1; i <= n; i++ {
		var j int
		fmt.Fscan(in, &j)
		f[j%m]++
		g[j]++
		if (g[j] & 1) == 0 {
			s[j%m]++
		}
	}
	j := 0
	for i := 1; i+i < m; i++ {
		j += cal(i, m-i)
	}
	if m%2 == 0 {
		j += f[m/2] >> 1
	}
	fmt.Println(j + f[0]/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
