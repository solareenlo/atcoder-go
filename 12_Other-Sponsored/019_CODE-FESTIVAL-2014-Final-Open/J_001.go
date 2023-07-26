package main

import "fmt"

func main() {
	var A, B, K int
	fmt.Scan(&A, &B, &K)
	if K == 0 {
		fmt.Println(1)
		return
	}
	if K == 1 {
		if A == B {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
		return
	}
	if A > B {
		A, B = B, A
	}
	d := gcd(A, B)
	A /= d
	B /= d
	lb := 0
	ub := B + 1
	for ub-lb > 1 {
		mid := (lb + ub) / 2
		k := (mid + int(float64(mid-1)*float64(A)/float64(B))) * 2
		if k <= K {
			lb = mid
		} else {
			ub = mid
		}
	}
	a := lb
	lb = 0
	ub = B + 1
	for ub-lb > 1 {
		mid := (lb + ub) / 2
		k := (mid - 1 + int(float64(mid-1)*float64(A)/float64(B))) * 2
		if k <= K {
			lb = mid
		} else {
			ub = mid
		}
	}
	b := lb
	ans := 1 + a + b
	if ans <= B+1 {
		fmt.Println(ans)
	} else {
		fmt.Println(B + 1)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
