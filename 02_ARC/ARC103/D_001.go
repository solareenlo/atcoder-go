package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	x, y, p := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
		p[i] = (abs(x[i]) + abs(y[i])) % 2
		if abs(p[i]) != p[0] {
			fmt.Println(-1)
			return
		}
	}

	m := 32
	k := m
	if abs(x[0]+y[0])%2 == 0 {
		k++
	}
	fmt.Println(k)
	for i := 0; i < m; i++ {
		fmt.Print(1<<i, " ")
		if i == k-1 {
			break
		}
	}
	if k-m != 0 {
		fmt.Print(1)
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		A := x[i] + y[i]
		B := x[i] - y[i]
		C := (A - 1 + (1 << m)) / 2
		D := (B - 1 + (1 << m)) / 2
		for j := 0; j < m; j++ {
			if (C>>j)&1 != 0 {
				if (D>>j)&1 != 0 {
					fmt.Print("R")
				} else {
					fmt.Print("U")
				}
			} else if (D>>j)&1 != 0 {
				fmt.Print("D")
			} else {
				fmt.Print("L")
			}
		}
		if k-m != 0 {
			fmt.Print("R")
		}
		fmt.Println()
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
