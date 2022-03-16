package main

import "fmt"

var k int

func solve(w, h int) int {
	res := 0
	for d := 0; d < h; d++ {
		for a := 1; a < w; a++ {
			R := 2*k - a*d + gcd(d, w) - 2
			if R < 0 {
				break
			}
			b := min(h-1-d, R/w)
			tmp := b
			b++
			if b+d < h && w*b-gcd(w-a, b+d)-gcd(a, b) <= R {
				tmp++
			}
			if d == 0 {
				res += tmp
			} else {
				res += tmp * 2
			}
		}
	}
	return res * 2
}

func main() {
	var w, h int
	fmt.Scan(&w, &h, &k)

	fmt.Println(solve(w, h) + solve(h, w))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
