package main

import "fmt"

func main() {
	var n, x, d int
	fmt.Scan(&n, &x, &d)

	if d == 0 {
		if x == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(n + 1)
		}
		return
	}

	p := d / abs(gcd(x, d))
	if d < 0 {
		x = -x
		d = -d
	}

	res := 0
	for i := 0; i < n+1; i++ {
		l := i*(i-1) - n*(n-1)/2
		r := (n-1)*n/2 - (n-i)*(n-1-i)
		l1 := (i-p)*(i-p-1) - n*(n-1)/2 - 2*x*p/d
		r1 := n*(n-1)/2 - (n-i+p)*(n-i+p-1) - 2*x*p/d
		if r < l1 || l > r1 || i < p {
			res += (r-l)/2 + 1
		} else {
			if l <= l1 && l1 <= r {
				res += (l1 - l) / 2
			}
			if l <= r1 && r1 <= r {
				res += (r - r1) / 2
			}
		}
	}
	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
