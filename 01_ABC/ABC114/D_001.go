package main

import "fmt"

var prime = make(map[int]int)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		PrimeFactorization(i)
	}

	p75 := primeCnt(75)
	p25 := primeCnt(25)
	p15 := primeCnt(15)
	p5 := primeCnt(5)
	p3 := primeCnt(3)
	fmt.Println(p75 + p25*(p3-1) + p15*(p5-1) + p5*(p5-1)*(p3-2)/2)
}

func primeCnt(num int) int {
	cnt := 0
	for _, v := range prime {
		if v >= num-1 {
			cnt++
		}
	}
	return cnt
}

func PrimeFactorization(n int) map[int]int {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			prime[i]++
			n /= i
		}
	}
	if n != 1 {
		prime[n]++
	}
	return prime
}
