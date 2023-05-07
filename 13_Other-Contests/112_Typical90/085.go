package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var k int
	fmt.Fscan(in, &k)

	vec := make([]int, 0)
	for i := 1; i*i <= k; i++ {
		if k%i == 0 {
			vec = append(vec, i)
			if i != k/i {
				vec = append(vec, k/i)
			}
		}
	}
	sort.Ints(vec)

	ans := 0
	for i := 0; i < len(vec); i++ {
		for j := i; j < len(vec); j++ {
			a := vec[i]
			b := vec[j]
			if k/a < b {
				continue
			}
			if k%(a*b) == 0 {
				c := k / (a * b)
				if b <= c {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
