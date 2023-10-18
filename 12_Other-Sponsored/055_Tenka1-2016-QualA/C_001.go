package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var r [26][26]bool

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var A, B string
		fmt.Fscan(in, &A, &B)
		P := false
		for j := 0; j < min(len(A), len(B)); j++ {
			if A[j] > B[j] {
				r[A[j]-'a'][B[j]-'a'] = true
				P = true
				break
			}
			if A[j] < B[j] {
				r[A[j]-'a'][B[j]-'a'] = true
				P = true
				break
			}
		}
		if P == false {
			if len(A) > len(B) {
				fmt.Println(-1)
				return
			}
		}
	}
	S := ""
	var used [26]bool
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if used[j] == true {
				continue
			}
			OK := true
			for k := 0; k < 26; k++ {
				if r[k][j] == true && used[k] == false {
					OK = false
				}
			}
			if OK == true {
				used[j] = true
				S += string('a' + j)
				break
			}
		}
		if i+1 != len(S) {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(S)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
