package main

import "fmt"

func main() {
	var n int
	var aa, ab, ba, bb string
	fmt.Scan(&n, &aa, &ab, &ba, &bb)

	if n == 2 || n == 3 || (ab == "A" && aa == "A") || (ab == "B" && bb == "B") {
		fmt.Println(1)
		return
	}

	f := make([]int, n)
	const mod = 1_000_000_007
	if ba == ab {
		f[0] = 1
		f[1] = 2
		for i := 2; i <= n-3; i++ {
			f[i] = (f[i-1] + f[i-2]) % mod
		}
		fmt.Println(f[n-3])
	} else {
		f[0] = 1
		for i := 1; i <= n-3; i++ {
			f[i] = f[i-1] * 2 % mod
		}
		fmt.Println(f[n-3])
	}
}
