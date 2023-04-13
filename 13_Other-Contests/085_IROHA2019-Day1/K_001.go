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
	mc := 1
	const MOD = 1000000007
	var ans int
	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(in, &m)
		tmp := 0
		for j := 0; j < m; j++ {
			var a int
			fmt.Fscan(in, &a)
			tmp += a * mc % MOD
			var t int
			for t = 1; a > 0; a, t = a/10, t*10 {

			}
			tmp += t % MOD * ans % MOD
		}
		mc = mc * m % MOD
		ans = tmp % MOD
	}
	fmt.Println(ans)
}
