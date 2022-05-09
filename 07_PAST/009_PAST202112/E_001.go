package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscan(in, &n)

	ans := 500
	for i := 1; i < len(n); i++ {
		a := int(n[i] - '0')
		b := int(n[i-1] - '0')
		tmp0 := 0
		if !((a-1)*(a-5) <= 0) {
			tmp0 = 1
		}
		tmp1 := 0
		if (b-1)*(b-5) <= 0 {
			tmp1 = 1
		}
		if a == b {
			ans += 301
		} else if tmp0^tmp1 != 0 {
			ans += 210
		} else {
			ans += 100
		}
	}
	fmt.Println(ans)
}
