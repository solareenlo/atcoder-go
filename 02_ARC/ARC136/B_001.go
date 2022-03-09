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

	s1 := make([]int, 5005)
	a := make([]int, 5005)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		s1[a[i]]++
	}

	s2 := make([]int, 5005)
	b := make([]int, 5005)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		s2[b[i]]++
	}

	for i := 1; i <= 5000; i++ {
		if s1[i]^s2[i] != 0 {
			fmt.Println("No")
			return
		}
	}

	for i := 1; i <= 5000; i++ {
		if s1[i] > 1 {
			fmt.Println("Yes")
			return
		}
	}

	S := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			tmp1 := 0
			if a[i] > a[j] {
				tmp1 = 1
			}
			tmp2 := 0
			if b[i] > b[j] {
				tmp2 = 1
			}
			S ^= (tmp1 ^ tmp2)
		}
	}
	if S != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
