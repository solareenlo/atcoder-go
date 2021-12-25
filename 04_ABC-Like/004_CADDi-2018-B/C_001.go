package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	res := 1
	if n == 1 {
		res = p
	} else {
		prime := PrimeFactorization(p)
		for k, v := range prime {
			if v >= n {
				res *= pow(k, v/n)
			}
		}
	}
	fmt.Println(res)
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}

func PrimeFactorization(n int) map[int]int {
	res := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			res[i]++
			n /= i
		}
	}
	if n != 1 {
		res[n]++
	}
	return res
}
