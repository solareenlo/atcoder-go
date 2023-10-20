package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a string
	fmt.Fscan(in, &a)
	b := a
	a = reverseString(a)
	A := strings.Split(a, "")

	for i := range A {
		if A[i] == "b" {
			A[i] = "d"
		} else if A[i] == "d" {
			A[i] = "b"
		} else if A[i] == "p" {
			A[i] = "q"
		} else if A[i] == "q" {
			A[i] = "p"
		}
	}

	if strings.Join(A, "") == b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
