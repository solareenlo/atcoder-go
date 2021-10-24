package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	s := make([]int, 1<<n)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	p := make([]int, 0)
	p = append(p, s[0])
	for i := 0; i < n; i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(p)))
		k := 0
		m := len(p)
		for j := 0; j < 1<<n; j++ {
			if s[j] == 0 {
				continue
			}
			if k == m {
				break
			}
			if p[k] > s[j] {
				p = append(p, s[j])
				s[j] = 0
				k++
			}
		}
		if k != m {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
