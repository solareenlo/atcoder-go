package main

import (
	"bufio"
	"fmt"
	"os"
)

func bit(x, y int) int {
	return (x >> y) & 1
}

func solve(a []int, s int) int {
	n := len(a)
	memo := make(map[int]int)
	for i := 0; i < 1<<(n/2); i++ {
		t := 0
		for k := 0; k < n/2; k++ {
			if bit(i, k) != 0 {
				t += a[k]
			} else {
				t += -a[k]
			}
		}
		memo[t] = i
	}
	for i := 0; i < 1<<(n-n/2); i++ {
		t := 0
		for k := 0; k < n-n/2; k++ {
			if bit(i, k) != 0 {
				t += a[k+n/2]
			} else {
				t += -a[k+n/2]
			}
		}
		if _, ok := memo[s-t]; ok {
			return i<<(n/2) | memo[s-t]
		}
	}
	return -1
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	a := make([]int, 0)
	b := make([]int, 0)
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		if (i % 2) != 0 {
			a = append(a, t)
		} else {
			b = append(b, t)
		}
	}

	retx := solve(a, x)
	rety := solve(b, y)
	if retx == -1 || rety == -1 {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
	dir := 0
	var ans string
	for i := 0; i < n; i++ {
		var new_dir int
		if i%2 != 0 {
			new_dir = 2 - 2*bit(retx, i/2)
		} else {
			new_dir = 3 - 2*bit(rety, i/2)
		}
		if (dir+1)%4 == new_dir {
			ans += "L"
		} else {
			ans += "R"
		}
		dir = new_dir
	}
	fmt.Println(ans)
}
