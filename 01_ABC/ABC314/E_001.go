package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var c, s [110]int
	var a [110][110]int
	var f [110 * 2]float64

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i], &s[i])
		for j := 1; j <= s[i]; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	for i := m - 1; i >= 0; i-- {
		ans := 1e18
		f[i] = 1e18
		for j := 1; j <= n; j++ {
			sum := 0.0
			z := 0.0
			for k := 1; k <= s[j]; k++ {
				if a[j][k] != 0 {
					sum += f[i+a[j][k]]
				} else {
					z++
				}
			}
			ans = math.Min(ans, (sum+float64(c[j]*s[j]))/(float64(s[j])-z))
			f[i] = ans
		}
	}
	fmt.Println(f[0])
}
