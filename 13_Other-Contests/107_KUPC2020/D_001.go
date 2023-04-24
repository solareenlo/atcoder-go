package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var d int
	for d = 2; d < n; d++ {
		if n%d == 0 {
			break
		}
	}
	if d == n {
		fmt.Println("impossible")
		return
	}

	if n%2 == 0 {
		fmt.Println(n / 2)
		for i := 0; i < n/2; i++ {
			fmt.Printf("2 %d %d\n", 2*i+1, 2*n-2*i-1)
		}
		return
	}

	ans := make([][]int, d)
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			ans[i] = append(ans[i], (i+j)%d)
		}
		for j := 0; j < (n/d-d)/2; j++ {
			ans[i] = append(ans[i], i)
			ans[i] = append(ans[i], d-i-1)
		}
	}
	for i := 0; i < d; i++ {
		for j := 0; j < n/d; j++ {
			ans[i][j] = 2*j*d + 1 + 2*ans[i][j]
		}
	}

	fmt.Println(d)
	for i := 0; i < d; i++ {
		fmt.Printf("%d ", n/d)
		for j := 0; j < n/d; j++ {
			if j < n/d-1 {
				fmt.Printf("%d ", ans[i][j])
			} else {
				fmt.Println(ans[i][j])
			}
		}
	}
}
