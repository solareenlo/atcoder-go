package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var a [2020]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var b [2020]int
	for i := 1; i <= n; i++ {
		flag := true
		for j := 0; j < n; j++ {
			b[j] = a[j]
		}
		for j := 0; j < n; j++ {
			if b[j%i] != 0 {
				if b[j] != 0 && b[j] != b[j%i] {
					flag = false
				}
			} else if b[j] != 0 {
				b[j%i] = b[j]
			}
		}
		if flag {
			fmt.Println(i)
			return
		}
	}
}
