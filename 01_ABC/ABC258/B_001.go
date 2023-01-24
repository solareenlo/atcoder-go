package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]string, 10)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	dx := [8]int{0, 1, 0, -1, 1, 1, -1, -1}
	dy := [8]int{1, 0, -1, 0, 1, -1, 1, -1}
	var ans string
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			for r := 0; r < 8; r++ {
				now := ""
				for d := 0; d < N; d++ {
					i := x + dx[r]*d
					j := y + dy[r]*d
					i = (i%N + N) % N
					j = (j%N + N) % N
					now += string(A[i][j])
				}
				if ans < now {
					ans = now
				}
			}
		}
	}
	fmt.Println(ans)
}
