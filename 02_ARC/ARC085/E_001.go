package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	r := rand.New(rand.NewSource(99))

	var n int
	fmt.Fscan(in, &n)

	const N = 222
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := -1 << 60
	for k := 0; k < 50000; k++ {
		b := make([]int, N)
		copy(b, a)
		t := 0
		for i := n; i > 0; i-- {
			t = 0
			for j := i; j <= n; j += i {
				t += b[j]
			}
			flag1 := t <= 0
			flag2 := r.Int63()%10 < 1
			if flag1 && !flag2 || !flag1 && flag2 {
				for j := i; j <= n; j += i {
					b[j] = 0
				}
			}
		}
		t = 0
		for i := 1; i <= n; i++ {
			t += b[i]
		}
		ans = max(ans, t)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
