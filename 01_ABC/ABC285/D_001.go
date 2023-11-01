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
	f := true
	n := make(map[string]string)
	for i := 0; i < N; i++ {
		var S, T string
		fmt.Fscan(in, &S, &T)
		for {
			_, ok := n[T]
			if ok {
				T = n[T]
			} else {
				break
			}
		}
		if S == T {
			f = false
		} else {
			n[S] = T
		}
	}
	if f {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
