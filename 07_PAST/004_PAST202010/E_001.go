package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var T string
	fmt.Scan(&N, &T)
	S := strings.Split(T, "")

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if i != N-j-1 && S[i] != S[j] {
				s := S[i]
				S[i] = S[j]
				S[j] = s
				fmt.Println(strings.Join(S, ""))
				return
			}
		}
	}
	fmt.Println("None")
}
