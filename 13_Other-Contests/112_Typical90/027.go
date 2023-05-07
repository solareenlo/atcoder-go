package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	S := make([]string, N)
	T := make(map[string]int)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S[i])
		if T[S[i]] == 0 {
			fmt.Println(i + 1)
			T[S[i]] = 1
		}
	}
}
