package main

import (
	"fmt"
	"strings"
)

const mod = 998244353

var (
	M = map[string]int{}
)

func dp(s string) int {
	if len(s) == 0 {
		return 1
	}
	if _, ok := M[s]; ok {
		return M[s]
	}
	ret := int(s[0]-'0'+1) * dp(s[1:]) % mod
	var tmp []string
	for i := 1; i <= len(s); i++ {
		tmp = strings.Split(s[0:i], "")
		for j := 2; i*j <= len(s); j++ {
			for k := i * (j - 1); k < i*j; k++ {
				if s[k] == '0' {
					tmp[k%i] = "0"
				}
			}
			ret = (ret + dp(strings.Join(tmp, ""))*dp(s[i*j:])) % mod
		}
	}
	M[s] = ret
	return M[s]
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(dp(s))
}
