package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := ""
	S := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		ans = string(S[n%36]) + ans
		n /= 36
		if n <= 0 {
			break
		}
	}
	fmt.Println(ans)
}
