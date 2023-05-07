package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n int
	fmt.Fscan(in, &n)
	ans := 1
	for n > 0 {
		n--
		sum := 0
		for i := 0; i < 6; i++ {
			var a int
			fmt.Fscan(in, &a)
			sum = (sum + a) % MOD
		}
		ans = ans * sum % MOD
	}
	fmt.Println(ans)
}
