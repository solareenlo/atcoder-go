package main

import "fmt"

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)

	a := make([]int, m+1)
	b := make([]int, m+1)
	c := make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Scan(&a[i], &b[i], &c[i])
		c[i] = p - c[i]
	}

	d := make([]int, n+1)
	for i := range d {
		d[i] = 1 << 62
	}
	d[1] = 0

	for i := 1; i < n<<1+1; i++ {
		for j := 1; j < m+1; j++ {
			if d[a[j]] != 1<<62 && d[b[j]] > d[a[j]]+c[j] {
				if i < n {
					d[b[j]] = d[a[j]] + c[j]
				} else {
					d[b[j]] = -1 << 60
				}
			}
		}
	}

	if d[n] == -1<<60 {
		fmt.Println(-1)
	} else {
		if d[n] < 0 {
			fmt.Println(-d[n])
		} else {
			fmt.Println(0)
		}
	}
}
