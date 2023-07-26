package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &array[i])
	}

	check := make([]bool, N+3)
	for i := 0; i < N; i++ {
		m := gcd(array[(i+1)%N], array[((i-1)%N+N)%N])
		if array[i]%m == 0 {
			check[i] = true
		} else {
			check[i] = false
		}
	}
	for i := 0; i < min(3, N); i++ {
		check[N+i] = check[i]
	}
	ans := N
	for i := 0; i < min(3, N); i++ {
		tmp := 0
		for j := i; j < N+i; j++ {
			if !check[j] {
				tmp++
				j = j + 2
			}
		}
		ans = min(tmp, ans)
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
