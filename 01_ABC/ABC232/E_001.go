package main

import "fmt"

func main() {
	var h, w, k, x1, y1, x2, y2 int
	fmt.Scan(&h, &w, &k, &x1, &y1, &x2, &y2)

	const mod = 998244353

	a, b, c, d := 1, 0, 0, 0
	for i := 0; i < k; i++ {
		na := ((h-1)*b + (w-1)*c) % mod
		nb := ((w-1)*d + a + (h-2)*b) % mod
		nc := ((h-1)*d + a + (w-2)*c) % mod
		nd := ((h-2)*d + (w-2)*d + b + c) % mod
		a = na
		b = nb
		c = nc
		d = nd
	}

	if x2 == x1 && y2 == y1 {
		fmt.Println(a)
	} else if x1 == x2 {
		fmt.Println(c)
	} else if y1 == y2 {
		fmt.Println(b)
	} else {
		fmt.Println(d)
	}
}
