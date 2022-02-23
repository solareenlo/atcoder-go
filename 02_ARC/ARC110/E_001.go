package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	if s == strings.Repeat(string(s[0]), n) {
		fmt.Println(1)
		return
	}

	dyn := make([]int, 4)
	curXor := 0
	const mod = 1_000_000_007
	for _, ch := range s {
		curXor ^= int(ch-'A') + 1
		tmp := 0
		if curXor != 0 {
			tmp = 1
		}
		dyn[curXor] = tmp
		for i := 0; i < 4; i++ {
			if i != curXor {
				dyn[curXor] = (dyn[curXor] + dyn[i]) % mod
			}
		}
	}
	fmt.Println(dyn[curXor])
}
