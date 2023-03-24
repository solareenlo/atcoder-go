package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	var T string
	var Q []string
	Q = append(Q, S)

	for len(Q) > 0 {
		V := Q[0]
		Q = Q[1:]
		if len(V) == 1 {
			T += V
		}
		if len(V) >= 2 {
			dep, bo := 0, 0
			for i := len(V) - 2; i >= 0; i-- {
				if V[i] == '+' || V[i] == '-' || V[i] == '*' || V[i] == '/' {
					dep++
				} else {
					dep--
				}
				if dep == -1 {
					bo = i
					break
				}
			}
			U1 := V[0:bo]
			U2 := V[bo : len(V)-1]
			Q = append(Q, U1)
			Q = append(Q, U2)
			T += string(V[len(V)-1])
		}
	}
	T = reverse(T)
	fmt.Println(T)
}

func reverse(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteByte(s[i])
	}
	return result.String()
}
