package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	ans := 0
	for A != 0 && B != 0 {
		minn := int(1e9)
		GCD := gcd(A, B)
		A /= GCD
		B /= GCD
		n := abs(A - B)
		if n <= 1 {
			ans += min(A, B)
			break
		}
		minn = A % n
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				minn = min(minn, min(A%i, A%(n/i)))
			}
		}
		ans += minn
		A -= minn
		B -= minn
	}
	fmt.Println(ans)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
