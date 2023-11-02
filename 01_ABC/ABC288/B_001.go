package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, K int
	fmt.Fscan(in, &N, &K)
	S := make([]string, K)
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &S[i])
	}
	sort.Strings(S)
	for i := 0; i < K; i++ {
		fmt.Fprintln(out, S[i])
	}
}
