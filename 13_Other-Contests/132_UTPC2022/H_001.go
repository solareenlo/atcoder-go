package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353

	var S string
	fmt.Fscan(in, &S)
	N := len(S)
	mR := 0
	ans := 0
	t := 1
	inv10 := 299473306
	for i := 0; i < 10; i++ {
		for j := mR; j < N; j++ {
			if int(S[j]-'0') == i {
				t = t * inv10 % mod
				ans = (ans + int('9'-S[j])*t) % mod
				mR = j
			}
		}
	}
	fmt.Println(ans)
}
