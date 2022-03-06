package main

import (
	"bufio"
	"fmt"
	"os"
)

var a = [300300]int{}

func ju(x, y int) int {
	if a[y]-a[x] >= 0 {
		return (a[y] - a[x]) / (y - x)
	}
	return -((a[x]-a[y]-1)/(y-x) + 1)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	tot := 0
	d := make([]int, 300300)
	for i := 1; i <= n; i++ {
		for (tot > 1) && (ju(d[tot-1], d[tot]) >= ju(d[tot], i)) {
			tot--
		}
		tot++
		d[tot] = i
	}

	Ans := make([]int, 300300)
	for i := 1; i <= tot; i++ {
		Ans[d[i]] = a[d[i]]
	}

	for i := 1; i < tot; i++ {
		t1 := a[d[i+1]] - a[d[i]]
		t2 := d[i+1] - d[i]
		if t1 > 0 {
			for j := d[i] + 1; j < d[i+1]; j++ {
				tmp := 0
				if d[i+1]-j < t1%t2 {
					tmp = 1
				}
				Ans[j] = Ans[j-1] + t1/t2 + tmp
			}
		} else {
			t1 = -t1
			for j := d[i+1] - 1; j > d[i]; j-- {
				tmp := 0
				if j-d[i] < t1%t2 {
					tmp = 1
				}
				Ans[j] = Ans[j+1] + t1/t2 + tmp
			}
		}
	}

	Sum := 0
	for i := 1; i <= n; i++ {
		Sum += Ans[i]
	}
	fmt.Println(Sum)
}
