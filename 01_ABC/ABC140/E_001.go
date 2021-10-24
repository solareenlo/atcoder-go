package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	f := make([]int, 200001)
	l := make([]int, 200001)
	r := make([]int, 200001)

	for i := 1; i < n+1; i++ {
		var p int
		fmt.Scan(&p)
		f[p] = i
		l[i], r[i] = i-1, i+1
	}

	res := 0
	for i := 1; i < n+1; i++ {
		x, y := l[f[i]], r[f[i]]
		if x > 0 {
			res += i * (x - l[x]) * (y - f[i])
		}
		if y <= n {
			res += i * (r[y] - y) * (f[i] - x)
		}
		l[y], r[x] = x, y
	}

	fmt.Println(res)
}
