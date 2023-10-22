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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	l := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			l = append(l, 1)
			if j < n-1 && a[j] > a[j+1] {
				l = append(l, n-1)
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(len(l))
	for _, i := range l {
		fmt.Println(i)
	}
}
