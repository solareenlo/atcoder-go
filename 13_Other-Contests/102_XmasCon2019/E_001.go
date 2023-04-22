package main

import (
	"fmt"
	"math"
)

const B = 320000

var isnt [B + 1]bool
var p int
var pr, f1, f2 [B]int
var m int

func main() {
	var n int
	fmt.Scan(&n)

	for i := 2; i <= B; i++ {
		if !isnt[i] {
			p++
			pr[p] = i
		}
		for j := 1; j <= p && pr[j]*i <= B; j++ {
			isnt[pr[j]*i] = true
			if i%pr[j] == 0 {
				break
			}
		}
	}

	ans := calc(n)
	for i := 1; pr[i] <= n/m; i++ {
		x := n
		for x != 0 {
			x = x / pr[i]
			ans += x
		}
	}
	fmt.Println(ans)
}

func calc(n int) int {
	m = int(math.Floor(math.Sqrt(float64(n))))
	for i := 1; i <= m; i++ {
		f1[i] = i - 1
		f2[i] = n/i - 1
	}
	for j := 1; pr[j] <= m; j++ {
		for i := 1; i <= m/pr[j]; i++ {
			f2[i] -= f2[i*pr[j]] - f1[pr[j]-1]
		}
		for i := m/pr[j] + 1; i <= n/pr[j]/pr[j] && i <= m; i++ {
			f2[i] -= f1[n/i/pr[j]] - f1[pr[j]-1]
		}
		for i := m; i/pr[j] >= pr[j]; i-- {
			f1[i] -= f1[i/pr[j]] - f1[pr[j]-1]
		}
	}
	res := 0
	for i := 1; i < m; i++ {
		res += i * (f2[i] - f2[i+1])
	}
	return res
}
