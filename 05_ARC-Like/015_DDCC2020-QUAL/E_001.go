package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

var N int

func ask(id int) bool {
	fmt.Print("?")
	for i := 0; i < N; i++ {
		fmt.Print(" ", id+i)
	}
	fmt.Println()
	var s string
	fmt.Fscan(in, &s)
	return s == "Red"
}

func main() {
	fmt.Fscan(in, &N)
	L := 1
	R := N + 1
	l := ask(1)

	for R-L > 1 {
		M := (L + R) / 2
		f := ask(M)
		if l == f {
			L = M
		} else {
			R = M
		}
	}

	ans := strings.Split(strings.Repeat(" ", 2*N), "")
	if l {
		ans[L-1] = "R"
		ans[L+N-1] = "B"
	} else {
		ans[L-1] = "B"
		ans[L+N-1] = "R"
	}

	for i := 1; i <= 2*N; i++ {
		if i == L || i == L+N {
			continue
		}
		if L < i && i < L+N {
			fmt.Print("? ", i)
			for j := 1; j < L; j++ {
				fmt.Print(" ", j)
			}
			for j := L + N + 1; j <= 2*N; j++ {
				fmt.Print(" ", j)
			}
			fmt.Println()
		} else {
			fmt.Print("? ", i)
			for j := L + 1; j < L+N; j++ {
				fmt.Print(" ", j)
			}
			fmt.Println()
		}
		var s string
		fmt.Fscan(in, &s)
		if s == "Red" {
			ans[i-1] = "R"
		} else {
			ans[i-1] = "B"
		}
	}
	fmt.Println("!", strings.Join(ans, ""))
}
