package main

import (
	"bytes"
	"fmt"
)

const Mod = int(1e9 + 7)

var (
	memo []map[string]int
	n    int
)

func ok(last4 string) bool {
	for i := 0; i < 4; i++ {
		t := []byte(last4)
		if i >= 1 {
			t[i-1], t[i] = t[i], t[i-1]
		}
		if bytes.Contains(t, []byte("AGC")) {
			return false
		}
	}
	return true
}

func dfs(cur int, last3 string) int {
	if v, ok := memo[cur][last3]; ok {
		return v
	}
	if cur == n {
		return 1
	}
	res := 0
	for _, c := range []string{"A", "C", "G", "T"} {
		if ok(last3 + c) {
			res = (res + dfs(cur+1, last3[1:]+c)) % Mod
		}
	}
	memo[cur][last3] = res
	return res
}

func main() {
	fmt.Scan(&n)
	memo = make([]map[string]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make(map[string]int)
	}
	fmt.Println(dfs(0, "TTT"))
}
