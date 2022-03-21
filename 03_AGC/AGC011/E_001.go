package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	n := len(s)
	a := [600001]int{}
	for i := 0; i < n; i++ {
		a[n-i] = int(s[i]-'0') * 9
	}

	var i int
	for i = 1; i <= n; i++ {
		if a[i] > 9 {
			a[i+1] += a[i] / 10
			a[i] %= 10
		}
	}

	if a[i] != 0 {
		n++
	}

	t := 0
	for i = 1; i <= n; i++ {
		t += a[i]
	}

	for i = 1; i <= n; i++ {
		a[1] += 9
		t += 9
		p := 1
		for a[p] > 9 {
			a[p+1]++
			a[p] -= 10
			t -= 9
			p++
		}
		if p > n {
			n++
		}
		if t <= 9*i {
			break
		}
	}
	fmt.Println(i)
}
