package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	var S, T string
	for i := 0; i < N; i++ {
		S += "01"
	}
	for i := 0; i < 2000; i++ {
		c := S[len(S)-1]
		S = S[:len(S)-1]
		fmt.Fprintln(out, "?", S+T)
		out.Flush()
		var res string
		fmt.Fscan(in, &res)
		if res != "Yes" {
			T = string(c) + T
		}
		if len(S)+len(T) == N {
			T = S + T
		}
		if len(T) >= N {
			break
		}
	}
	fmt.Fprintln(out, "!", T)
	out.Flush()
}
