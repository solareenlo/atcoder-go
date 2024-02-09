package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	var bs, msks [105]int
	for i := 0; i < m; i++ {
		var ki int
		fmt.Fscan(in, &ki)
		for j := 0; j < ki; j++ {
			var aj, bj int
			fmt.Fscan(in, &aj, &bj)
			aj--
			bs[i] |= bj << aj
			msks[i] |= 1 << aj
		}
	}

	nbits := 1 << n
	for bits := 0; bits < nbits; bits++ {
		ok := true
		for i := 0; ok && i < m; i++ {
			ok = (((bits & msks[i]) ^ bs[i]) != msks[i])
		}
		if ok {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
