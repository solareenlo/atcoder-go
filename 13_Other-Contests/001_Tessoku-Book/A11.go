package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, X int
	fmt.Fscan(in, &N, &X)
	vec := make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &vec[i])
	}
	left := 1
	right := N
	for left < right {
		mid := (left + right) / 2
		if X <= vec[mid] {
			right = mid
		} else if X > vec[mid] {
			left = mid + 1
		}
	}
	fmt.Println(right)
}
