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
	var a [2 << 17]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	M := make(map[int]int)
	M[0] = 1
	sum := 0
	var L [2 << 17]int
	for i := 0; i < n; i++ {
		sum += a[i]
		L[i+1] = L[i] + M[sum]
		M[sum]++
	}
	M = make(map[int]int)
	sum = 0
	M[0] = 1
	var R [2 << 17]int
	for i := n; i > 0; i-- {
		sum += a[i-1]
		R[i-1] = R[i] + M[sum]
		M[sum]++
	}
	ans := int(9e18)
	for i := 0; i < n; i++ {
		now := L[i] + R[i+1]
		if now < ans {
			ans = now
		}
	}
	fmt.Println(ans)
}
