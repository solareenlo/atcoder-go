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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	total := accumulate(a)
	ans := 0
	p10 := 1
	for p10 <= total {
		f := make([]int, 1)
		for _, x := range a {
			sz := len(f)
			resize(&f, 2*sz)
			for i := 0; i < sz; i++ {
				f[i+sz] = (f[i] + x) % (10 * p10)
			}
			sort.Ints(f)
			new_sz := 0
			for i := 0; i < 2*sz; i++ {
				if i > 0 && i < 2*sz-1 && f[i+1]-f[new_sz-1] <= p10 {
					continue
				}
				f[new_sz] = f[i]
				new_sz++
			}
			resize(&f, new_sz)
		}
		ans += f[len(f)-1] / p10
		p10 *= 10
	}
	fmt.Println(ans)
}

func accumulate(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}
