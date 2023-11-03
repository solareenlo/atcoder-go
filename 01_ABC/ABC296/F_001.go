package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 220000

	var a, b, frq, pos [N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		frq[a[i]]++
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		frq[b[i]]--
	}

	for i := 1; i <= n; i++ {
		if frq[i] != 0 {
			fmt.Println("No")
			return
		}
	}

	for i := 1; i <= n; i++ {
		frq[a[i]]++
	}

	for i := 1; i <= n; i++ {
		if frq[i] > 1 {
			fmt.Println("Yes")
			return
		}
	}

	for i := 1; i <= n; i++ {
		pos[a[i]] = i
	}
	count := 0
	for i := 1; i <= n; i++ {
		if pos[i] != i {
			count ^= 1
			pos[a[i]] = pos[i]
			a[i], a[pos[i]] = a[pos[i]], a[i]
			pos[i] = i
		}
	}

	for i := 1; i <= n; i++ {
		pos[b[i]] = i
	}
	for i := 1; i <= n; i++ {
		if pos[i] != i {
			count ^= 1
			pos[b[i]] = pos[i]
			b[i], b[pos[i]] = b[pos[i]], b[i]
			pos[i] = i
		}
	}

	if count != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
