package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	if a[0] != 1 {
		fmt.Println(0)
		return
	}
	tr := make([][]int, n)
	fo := make([][]int, n)
	for i := 0; i < n; i++ {
		tr[i] = make([]int, n)
		fo[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		tr[i][i] = 1
		fo[i][i] = 1
	}
	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			tr[i][j] = fo[i+1][j]
			fo[i][j] = tr[i][j]
			for k := i; k < j; k++ {
				if a[i] < a[k+1] {
					fo[i][j] = (fo[i][j] + tr[i][k]*fo[k+1][j]%MOD) % MOD
				}
			}
		}
	}
	fmt.Println(tr[0][n-1])
}
