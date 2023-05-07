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
	var s string
	fmt.Fscan(in, &n, &s)

	var table [8]int
	table[0] = 1
	target := "atcoder"
	for i := 0; i < n; i++ {
		for j := 0; j < 7; j++ {
			if s[i] == target[j] {
				table[j+1] = (table[j+1] + table[j]) % MOD
				break
			}
		}
	}

	fmt.Println(table[7])
}
