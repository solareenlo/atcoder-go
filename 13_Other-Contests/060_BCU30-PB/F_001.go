package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1e9 + 7

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, a int
	fmt.Fscan(in, &n, &a)
	dp1 := a
	dp2 := 1
	dp3 := 0
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a)
		dp3 = (dp1 + dp3*2) % mod
		dp1 = ((dp1 + dp2) * a) % mod
		dp2 = dp2 * 2 % mod
	}
	fmt.Println((dp1 + dp3) % mod)
}
