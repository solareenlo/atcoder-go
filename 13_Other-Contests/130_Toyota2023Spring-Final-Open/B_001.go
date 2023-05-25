package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	S := make([]int, n)
	P := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &S[i], &P[i])
	}
	ord := make([]int, n)
	for i := range ord {
		ord[i] = i
	}
	sort.Slice(ord, func(i, j int) bool {
		return P[ord[i]]*(100-P[ord[j]])*S[ord[i]]-P[ord[j]]*(100-P[ord[i]])*S[ord[j]] > 0
	})
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", ord[i]+1)
	}
}
