package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const t = 10000

	var n, x, y, a int
	fmt.Fscan(in, &n, &x, &y, &a)

	h := new(big.Int)
	s := new(big.Int)
	h.SetBit(h, t+a, 1)
	s.SetBit(s, t, 1)

	for i := 2; i <= n; i++ {
		var a uint
		fmt.Fscan(in, &a)
		if i&1 != 0 {
			h.Or(new(big.Int).Lsh(h, a), new(big.Int).Rsh(h, a))
		} else {
			s.Or(new(big.Int).Lsh(s, a), new(big.Int).Rsh(s, a))
		}
	}

	if h.Bit(t+x) != 0 && s.Bit(t+y) != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
