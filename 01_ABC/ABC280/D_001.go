package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)
	L := K

	ans := -1
	for i := 2; i*i <= L; i++ {
		if K%i != 0 {
			continue
		}
		cnt := 0
		for K%i == 0 {
			cnt++
			K /= i
		}

		p := i
		for cnt > 0 {
			q := p
			for q%i == 0 {
				cnt--
				q /= i
			}
			p += i
		}
		ans = max(ans, p-i)
	}
	fmt.Println(max(ans, K))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
