package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	S := make([]int, N)
	E := make([]int, N)
	L := -int(2e9)
	R := int(2e9)
	for i := 0; i < N; i++ {
		var h, m, s, ms int
		var hh, mm, ss, msms int
		fmt.Scanf("%d:%d:%d.%d %d:%d:%d.%d\n", &h, &m, &s, &ms, &hh, &mm, &ss, &msms)
		S[i] = (((h-21)*60+m)*60+s)*1000 + ms
		E[i] = (((hh-21)*60+mm)*60+ss)*1000 + msms
		if E[i] <= S[i] {
			L = max(L, S[i])
			R = min(R, E[i])
		}
	}

	for i := 0; i < N; i++ {
		if E[i] <= S[i] {
			fmt.Println(E[i] - S[i] + 1000)
		} else {
			if E[i] <= L-1000 || R+1000 <= S[i] {
				fmt.Println(E[i] - S[i])
			} else if S[i] <= L-1000 && R+1000 <= E[i] {
				fmt.Println(E[i] - S[i] + 1000)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
