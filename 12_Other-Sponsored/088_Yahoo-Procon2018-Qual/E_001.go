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
	var d [100000]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
	}
	var S [100010]int
	S[0] = 0
	for i := 0; i < n; i++ {
		S[i+1] = S[i] + d[i]
	}
	var T bool
	if S[n]%2 != 0 {
		T = true
	} else {
		T = false
	}
	if T {
		j := n - 1
		for j > 0 && d[j-1] == d[n-1] {
			j--
		}
		d[j]--
		for i := j + 1; i <= n; i++ {
			S[i]--
		}
	}
	F := true
	for k, j := 1, 0; k < n; k++ {
		for j < n-k && d[j] < k {
			j++
		}
		j = min(j, n-k)
		if S[n]-S[n-k]-k*(k-1) > S[j]+(n-k-j)*k {
			F = false
			break
		}
	}
	if !F {
		fmt.Println("ABSOLUTELY NO")
	} else if T {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
