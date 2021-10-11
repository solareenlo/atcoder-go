package main

import "fmt"

func main() {
	var l int
	fmt.Scan(&l)

	now, n := 1, 0
	for now <= l {
		now <<= 1
		n++
	}

	u, v, w := make([]int, 0), make([]int, 0), make([]int, 0)
	for i := 1; i < n; i++ {
		u = append(u, i)
		v = append(v, i+1)
		w = append(w, 1<<(i-1))
		u = append(u, i)
		v = append(v, i+1)
		w = append(w, 0)
	}

	for i := n - 1; i >= 1; i-- {
		if l-(1<<(i-1)) >= (1 << (n - 1)) {
			u = append(u, i)
			v = append(v, n)
			w = append(w, l-(1<<(i-1)))
			l -= 1 << (i - 1)
		}
	}

	fmt.Println(n, len(u))
	for i := 0; i < len(u); i++ {
		fmt.Println(u[i], v[i], w[i])
	}
}
