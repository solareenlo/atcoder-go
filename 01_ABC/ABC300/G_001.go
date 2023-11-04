package main

import (
	"fmt"
	"math/bits"
)

var primes [25]int = [25]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

const N = 1 << 18

var f [N][25]int

func cal(n, p int) int {
	if n <= 4 {
		if n >= 3 && p == 0 {
			return n - 1
		}
		return n
	}
	if p == 0 {
		return 64 - countLeadingZeros(uint64(n))
	}
	if n < N && f[n][p] != 0 {
		return f[n][p]
	}
	ans := cal(n, p-1)
	if n >= primes[p] {
		ans += cal(n/primes[p], p)
	}
	if n < N {
		f[n][p] = ans
	}
	return ans
}

var n, p int

func main() {
	fmt.Scan(&n, &p)
	idx := findIndex(len(primes), func(i int) bool {
		return primes[i] == p
	})
	fmt.Println(cal(n, idx))
}

func findIndex(lim int, predicate func(i int) bool) int {
	for i := 0; i < lim; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func countLeadingZeros(x uint64) int {
	return bits.LeadingZeros64(x)
}
