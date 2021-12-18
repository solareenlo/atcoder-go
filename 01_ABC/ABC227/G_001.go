package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	l, r := n-k+1, n
	t := make([]int, 1<<20)
	for i := 0; i < k; i++ {
		t[i] = l + i
	}

	mod := 998244353
	res := 1
	isp := make([]bool, 1<<20)
	for i := 2; i < 1<<20; i++ {
		if !isp[i] {
			d := i
			cnt := 0
			for d <= k {
				cnt -= k / d
				d *= i
			}
			nl := (l + i - 1) / i * i
			for nl <= r {
				id := nl - l
				for t[id]%i == 0 {
					t[id] /= i
					cnt++
				}
				nl += i
			}
			res = res * (cnt + 1) % mod
		}
		for j := i + i; j < 1<<20; j += i {
			isp[j] = true
		}

	}
	for i := 0; i < k; i++ {
		if t[i] > 1 {
			res = res * 2 % mod
		}
	}
	fmt.Println(res)
}
