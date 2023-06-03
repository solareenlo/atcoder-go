package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007
	const B = 1001

	var n, q int
	var s string
	fmt.Fscan(in, &n, &q, &s)
	rs := s
	rs = reverseString(rs)
	powB := make([]int, n+1)
	hash := make([]int, n+1)
	rhash := make([]int, n+1)
	powB[0] = 1
	for i := 1; i <= n; i++ {
		powB[i] = powB[i-1] * B % mod
	}
	hash[0] = 0
	for i := 1; i <= n; i++ {
		hash[i] = (hash[i-1]*B + (int(s[i-1]-'a') + 1)) % mod
	}
	rhash[0] = 0
	for i := 1; i <= n; i++ {
		rhash[i] = (rhash[i-1]*B + (int(rs[i-1]-'a') + 1)) % mod
	}
	for q > 0 {
		q--
		var l, r int
		fmt.Fscan(in, &l, &r)
		hash1 := ((hash[r]-powB[r-l+1]*hash[l-1])%mod + mod) % mod
		hash2 := ((rhash[n+1-l]-powB[r-l+1]*rhash[n-r])%mod + mod) % mod
		if hash1 == hash2 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
