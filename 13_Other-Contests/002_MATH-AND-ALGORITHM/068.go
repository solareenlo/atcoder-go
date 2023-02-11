package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	v := make([]int, k)
	for i := range v {
		fmt.Fscan(in, &v[i])
	}

	ans := 0
	for b := 1; b < (1 << k); b++ {
		d := 1
		for i := 0; i < k; i++ {
			if b>>i&1 != 0 {
				d = lcm(d, v[i])
			}
		}
		if bits.OnesCount(uint(b))%2 != 0 {
			ans += n / d
		} else {
			ans -= n / d
		}
	}

	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
