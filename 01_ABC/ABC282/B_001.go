package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	S := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S[i])
	}
	ans := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			ok := true
			for k := 0; k < M; k++ {
				if S[i][k] == 'x' && S[j][k] == 'x' {
					ok = false
				}
			}
			if ok {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
