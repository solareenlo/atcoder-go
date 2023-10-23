package main

import (
	"bufio"
	"fmt"
	"os"
)

func get(dep, ub int) int {
	if ub == 0 {
		return 0
	}
	if dep == 0 {
		return 1
	}
	if (dep & 1) == 0 {
		return get(dep/2, (ub+1)/2)
	}
	if ub%2 == 0 {
		return get(dep/2, (ub-2)/2)
	} else {
		return get(dep/2, (ub+1)/2)
	}
}

var d, l [303030]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscan(in, &num)
	sum := 0
	for i := 0; i < num; i++ {
		fmt.Fscan(in, &d[i], &l[i])
		sum += l[i]
	}
	ans := 0
	bef := 0
	for i := 0; i < num; i++ {
		t := get(sum/2, bef+l[i]) ^ get(sum/2, bef)
		if t != 0 {
			ans ^= d[i]
		}
		bef += l[i]
	}
	fmt.Println(ans)
}
