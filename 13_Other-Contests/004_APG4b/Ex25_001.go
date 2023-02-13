package main

import (
	"fmt"
	"math/big"
)

func main() {
	A := new(big.Int)
	B := new(big.Int)
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var x int
		fmt.Scan(&x)
		A = add(A, x)
	}
	var M int
	fmt.Scan(&M)
	for i := 0; i < M; i++ {
		var x int
		fmt.Scan(&x)
		B = add(B, x)
	}

	var com string
	fmt.Scan(&com)

	switch com {
	case "intersection":
		printSet(intersection(A, B))
	case "union_set":
		printSet(unionSet(A, B))
	case "symmetric_diff":
		printSet(symmetricDiff(A, B))
	case "subtract":
		var x int
		fmt.Scan(&x)
		printSet(subtract(A, x))
	case "increment":
		printSet(increment(A))
	case "decrement":
		printSet(decrement(A))
	}
}

func intersection(A, B *big.Int) *big.Int {
	return A.And(A, B)
}

func unionSet(A, B *big.Int) *big.Int {
	return A.Or(A, B)
}

func symmetricDiff(A, B *big.Int) *big.Int {
	return A.Xor(A, B)
}

func increment(A *big.Int) *big.Int {
	if A.Bit(49) != 0 {
		A = A.Lsh(A, 1)
		A.SetBit(A, 0, 1)
	} else {
		A = A.Lsh(A, 1)
	}
	return A
}

func decrement(A *big.Int) *big.Int {
	if A.Bit(0) != 0 {
		A = A.Rsh(A, 1)
		A.SetBit(A, 49, 1)
	} else {
		A = A.Rsh(A, 1)
	}
	return A
}

func subtract(A *big.Int, x int) *big.Int {
	A.SetBit(A, x, 0)
	return A
}

func add(A *big.Int, x int) *big.Int {
	A.SetBit(A, x, 1)
	return A
}

func printSet(S *big.Int) {
	cont := make([]int, 0)
	for i := 0; i < 50; i++ {
		if S.Bit(i) != 0 {
			cont = append(cont, i)
		}
	}
	for i := 0; i < len(cont); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(cont[i])
	}
	fmt.Println()
}
