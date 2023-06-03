package main

import "fmt"

func main() {
	var to [30][3 << 17]int

	var n, k int
	fmt.Scan(&n, &k)
	for i := 0; i < n+1; i++ {
		t := i
		s := i
		for ; t != 0; t /= 10 {
			s -= t % 10
		}
		to[0][i] = s
	}
	for j := 0; j < 29; j++ {
		for i := 0; i < n+1; i++ {
			to[j+1][i] = to[j][to[j][i]]
		}
	}
	for i := 1; i <= n; i++ {
		v := i
		for j := 0; j < 29; j++ {
			if ((k >> j) & 1) != 0 {
				v = to[j][v]
			}
		}
		fmt.Println(v)
	}
}
