package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	A := [50][50]bool{}
	for i := 0; i < N; i++ {
		var k int
		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			var a int
			fmt.Scan(&a)
			A[i][a-1] = true
		}
	}

	var P, Q int
	fmt.Scan(&P, &Q)
	B := [50]int{}
	for i := 0; i < P; i++ {
		fmt.Scan(&B[i])
	}

	ans := 0
	for i := 0; i < N; i++ {
		now := 0
		for j := 0; j < P; j++ {
			if A[i][B[j]-1] {
				now++
			}
		}
		if now >= Q {
			ans++
		}
	}
	fmt.Println(ans)
}
