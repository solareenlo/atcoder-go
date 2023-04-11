package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var s string
	var d int
	fmt.Fscan(in, &s, &d)
	ac := 0
	var dp [105]int
	for i := 0; i < len(s); i++ {
		var tmp [105]int
		for j := 0; j < d; j++ {
			for k := 0; k < 10; k++ {
				tmp[(j+k)%d] = (tmp[(j+k)%d] + dp[j]) % mod
			}
		}
		for j := 0; j < int(s[i]-'0'); j++ {
			tmp[(ac+j)%d] = (tmp[(ac+j)%d] + 1) % mod
		}
		for j := 0; j < d; j++ {
			dp[j] = tmp[j]
		}
		ac += int(s[i] - '0')
	}
	if ac%d == 0 {
		fmt.Println((dp[0] + 1 - 1 + mod) % mod)
	} else {
		fmt.Println((dp[0] - 1 + mod) % mod)
	}
}
