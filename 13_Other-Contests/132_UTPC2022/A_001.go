package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var K int
		fmt.Fscan(in, &K)
		S := strconv.Itoa(K)
		ok := true
		for i := 0; i < len(S); i++ {
			if S[i] != S[0] {
				ok = false
			}
		}
		if ok {
			fmt.Println(S)
		} else if K%3 == 0 && K/3 < 10 {
			for i := 0; i < 3; i++ {
				fmt.Print(K / 3)
			}
			fmt.Println(0)
		} else if K%9 == 0 && K/9 < 10 {
			for i := 0; i < 9; i++ {
				fmt.Print(K / 9)
			}
			fmt.Println(0)
		} else {
			fmt.Println(-1)
		}
	}
	return
}
