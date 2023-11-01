package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var S string
	fmt.Scan(&S)
	cnt := 0
	for i := 1; i < N; i++ {
		cnt = 0
		for j := 0; j+i+1 <= N; j++ {
			if S[j] != S[j+i] {
				cnt++
			} else {
				break
			}
		}
		fmt.Println(cnt)
	}
}
