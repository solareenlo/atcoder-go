package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	z := 150
	ans := 0
	var b [40000000]int
	for k := -z; k <= z; k++ {
		for i := 0; i < n; i++ {
			b[a[i]-i*k+20000000]++
		}
		for i := 0; i < n; i++ {
			if b[a[i]-i*k+20000000] != 0 {
				ans += b[a[i]-i*k+20000000] * (b[a[i]-i*k+20000000] - 1) / 2
				b[a[i]-i*k+20000000] = 0
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if z*(j-i) > 100000 {
				break
			}
			if abs(a[i]-a[j])%(j-i) == 0 && abs(a[i]-a[j])/(j-i) > z {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
