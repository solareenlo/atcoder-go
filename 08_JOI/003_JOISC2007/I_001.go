package main

import (
	"bufio"
	"fmt"
	"os"
)

var UF [10000]int

func FIND(a int) int {
	if UF[a] < 0 {
		return a
	}
	UF[a] = FIND(UF[a])
	return UF[a]
}

func UNION(a, b int) {
	a = FIND(a)
	b = FIND(b)
	if a == b {
		return
	}
	UF[a] += UF[b]
	UF[b] = a
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)
	for i := 0; i < a; i++ {
		UF[i] = -1
	}
	for i := 0; i < b; i++ {
		var c, d int
		fmt.Fscan(in, &c, &d)
		UNION(c-1, d-1)
	}
	ret := 0
	for i := 0; i < a; i++ {
		if UF[i] < 0 {
			ret++
		}
	}
	fmt.Println(ret - 1)
}
