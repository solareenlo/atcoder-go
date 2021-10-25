package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	n := gcd(a, b)

	cnt := 1
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			cnt++
			for n%i == 0 {
				n /= i
			}
		}
	}

	if n != 1 {
		cnt++
	}
	fmt.Println(cnt)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
