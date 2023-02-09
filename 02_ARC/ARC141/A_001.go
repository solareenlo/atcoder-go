package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)
	a := make([]int, 110)
	for T > 0 {
		T--
		var n int
		fmt.Fscan(in, &n)
		if n < 100 {
			fmt.Fprintln(out, n/11*11)
			continue
		}
		m := n
		t := 1
		c := 1
		for m > 9 {
			m /= 10
			t *= 10
			c++
		}
		ans := t - 1
		pw := 1
		for i := 1; i < c; i++ {
			pw *= 10
			if c%i != 0 {
				continue
			}
			m := n
			v := 0
			res := 0
			for j := 1; j <= c/i; j++ {
				a[j] = m % pw
				m /= pw
			}
			v = a[c/i]
			for j := c / i; j >= 1; j-- {
				if a[j] != v {
					if a[j] < v {
						v--
					}
					break
				}
			}
			if v < pw/10 {
				continue
			}
			for j := 1; j <= c/i; j++ {
				res = res*pw + v
			}
			ans = max(ans, res)
		}
		fmt.Fprintln(out, ans)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
