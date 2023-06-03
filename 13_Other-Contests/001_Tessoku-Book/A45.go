package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	m := make(map[string]int)
	m["W"] = 0
	m["B"] = 1
	m["R"] = 2
	var N int
	var c string
	fmt.Fscan(in, &N, &c)
	now := 3 - m[c]
	var a string
	fmt.Fscan(in, &a)
	for i := 0; i < N; i++ {
		now += m[string(a[i])]
	}
	if now%3 != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
