package main

import "fmt"

func main() {

	const MAX = 2000000

	prime := make([]bool, MAX)
	for i := range prime {
		prime[i] = true
	}
	for i := 2; i < MAX; i++ {
		if prime[i] {
			for j := i * 2; j < MAX; j += i {
				prime[j] = false
			}
		}
	}
	A := make([]int, 0)
	B := make([]int, 0)
	for i := 10000; i < MAX; i++ {
		if prime[i] {
			A = append(A, i)
			A, B = B, A
		}
	}
	for i := 1; i <= MAX; i++ {
		if len(A) < 100000 {
			A = append(A, i)
		}
		if len(B) < 100000 {
			B = append(B, i)
		}
	}
	for i := 0; i < 100000; i++ {
		fmt.Print(A[i])
		if i < 99999 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for i := 0; i < 100000; i++ {
		fmt.Print(B[i])
		if i < 99999 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
