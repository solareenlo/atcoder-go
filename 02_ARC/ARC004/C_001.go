package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d/%d", &a, &b)

	flag := true
	n := a * 2 / b
	for i := 0; i < 2; i++ {
		if ((b*(n+1)-2*a)*n)%(2*b) == 0 {
			m := ((b*(n+1) - 2*a) * n) / (2 * b)
			if m >= 1 && m <= n {
				fmt.Println(n, m)
				flag = false
			}
		}
		n++
	}

	if flag {
		fmt.Println("Impossible")
	}
}
