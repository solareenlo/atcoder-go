package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	S = "0" + S
	ans := 0
	N := len(S)
	sub := 0

	ren := 0
	for i := N - 2; i >= 0; i-- {
		if S[i] == '2' && S[i+1] == '5' {
			ren++
			sub += ren
		} else if S[i] == '5' && S[i-1] == '2' {
			sub = sub
		} else {
			ans += sub
			ren = 0
			sub = 0
		}
	}
	fmt.Println(ans)
}
