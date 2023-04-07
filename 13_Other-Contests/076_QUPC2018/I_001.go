package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)
	M := int(1e6)
	num := make([]int, M+1)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		num[a]++
	}

	ans := 0
	sum := make([]int, M+1)
	for i := M; i > 0; i-- {
		r := M / i * i
		s := 0
		for l := i; l <= M; l += i {
			for r > 0 && l+r >= K {
				s += num[r]
				r -= i
			}
			sum[i] += s * num[l]
		}
		for j := 2 * i; j <= M; j += i {
			sum[i] -= sum[j]
		}
		if K%i == 0 {
			ans += sum[i]
			if 2*i >= K {
				ans -= num[i]
			}
		}
	}
	ans /= 2
	fmt.Println(ans)
}
