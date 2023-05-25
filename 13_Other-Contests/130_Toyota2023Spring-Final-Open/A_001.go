package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, v int
	fmt.Fscan(in, &n, &m, &v)
	v *= 2
	ans := 0
	for y := 1; y <= n; y++ {
		if v%y != 0 {
			continue
		}
		for z := 1; z <= m; z++ {
			if (v/y)%z != 0 {
				continue
			}
			dx := v/y/z - (y-1)*m - (z - 1)
			if dx <= 0 {
				continue
			}
			if dx&1 != 0 {
				continue
			}
			dx >>= 1
			if (dx-1)/m+y <= n && (dx-1)%m+z <= m {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
