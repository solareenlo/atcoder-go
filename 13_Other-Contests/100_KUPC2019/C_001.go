package main

import "fmt"

func main() {
	var M, K int
	fmt.Scan(&M, &K)

	ans := 0
	cnt := 1
	for cnt < 2*M+1 {
		cnt *= 2*K + 1
		ans += 1
	}
	fmt.Println(ans)
}
