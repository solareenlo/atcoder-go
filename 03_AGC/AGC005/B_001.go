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

	b := [200100][2]int{}
	b[0][0] = 0
	j := 0
	a := [200100]int{}
	k := 0
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		s := 0
		for a[i] <= b[j][0] {
			s = s + b[j][1]
			k = k - b[j][0]*b[j][1]
			j--
		}
		j++
		b[j][0] = a[i]
		s++
		b[j][1] = s
		k = k + (a[i] * s)
		sum = sum + k
	}
	fmt.Println(sum)
}
