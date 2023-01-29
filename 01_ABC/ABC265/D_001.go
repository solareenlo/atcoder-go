package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, p, q, r int
	fmt.Fscan(in, &n, &p, &q, &r)

	var s [200005]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		s[i] = s[i-1] + s[i]
	}

	y, z, w := 2, 3, 4
	for x := 1; x <= n-2; x++ {
		for s[y-1]-s[x-1] < p && y <= n {
			y++
		}
		for s[z-1]-s[y-1] < q && z <= n {
			z++
		}
		for s[w-1]-s[z-1] < r && w <= n {
			w++
		}
		if s[w-1]-s[z-1] != r || s[z-1]-s[y-1] != q || s[y-1]-s[x-1] != p {
			continue
		}
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
