package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	A := [100][100]bool{}
	for i := 0; i < q; i++ {
		var t int
		fmt.Scan(&t)
		if t == 1 {
			var a, b int
			fmt.Scan(&a, &b)
			a--
			b--
			A[a][b] = true
		} else if t == 2 {
			var a int
			fmt.Scan(&a)
			a--
			for i := 0; i < n; i++ {
				if A[i][a] {
					A[a][i] = true
				}
			}
		} else {
			var a int
			fmt.Scan(&a)
			a--
			A2 := [100]bool{}
			for i := 0; i < n; i++ {
				if A[a][i] {
					for j := 0; j < n; j++ {
						if j != a && A[i][j] {
							A2[j] = true
						}
					}
				}
			}
			for i := 0; i < n; i++ {
				if A2[i] {
					A[a][i] = true
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] {
				fmt.Print("Y")
			} else {
				fmt.Print("N")
			}
		}
		fmt.Println()
	}
}
