package main

import "fmt"

func main() {
	var k, n int
	fmt.Scan(&k, &n)

	const N = 100000
	s := make([]int, N+10)
	for i := 1; i <= N; i++ {
		s[i] = s[i/10] + i%10
	}

	p := map[int]int{}
	for i := 0; i < N; i++ {
		p[(s[i]%k-i%k+k)%k]++
	}

	t := 0
	ans := 0
	for t = 0; (t+1)*N-1 <= n; t++ {
		ans += p[(t*N%k-s[t]%k+k)%k]
	}
	for i := 0; i+t*N <= n; i++ {
		if (s[i]+s[t])%k == (i+t*N)%k {
			ans++
		}
	}
	fmt.Println(ans - 1)
}
