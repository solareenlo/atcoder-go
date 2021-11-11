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

	cnt := make([]int, 1000010)
	d := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		cnt[a]++
		d = gcd(d, a)
	}
	if d > 1 {
		fmt.Println("not coprime")
		return
	}

	for i := 2; i < 1000010; i++ {
		sum := 0
		for j := i; j < 1000010; j += i {
			sum += cnt[j]
		}
		if sum > 1 {
			fmt.Println("setwise coprime")
			return
		}
	}

	fmt.Println("pairwise coprime")
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
