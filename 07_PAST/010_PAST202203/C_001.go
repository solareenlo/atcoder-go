package main

import "fmt"

func main() {
	T := "abcdefghijklmnopqrstuvwxyz"
	var S string
	fmt.Scan(&S)
	N := len(S)
	A := (N - 1) / 3
	B := N % 3
	if B == 0 {
		B = 3
	}
	for i := 0; i < B; i++ {
		fmt.Print(string(S[i]))
	}
	fmt.Println(string(T[A-1]))
}
