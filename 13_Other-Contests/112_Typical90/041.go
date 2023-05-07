package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type point struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]point, n+2)
	for x := 1; x <= n; x++ {
		fmt.Fscan(in, &a[x].x, &a[x].y)
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(a, b int) bool {
		if tmp[a].x == tmp[b].x {
			return tmp[a].y < tmp[b].y
		}
		return tmp[a].x < tmp[b].x
	})

	var res [500001]point
	cnt := 0
	for x := 1; x <= n; x++ {
		for cnt >= 2 && xmul(res[cnt-1], res[cnt], a[x]) < 0 {
			cnt--
		}
		cnt++
		res[cnt] = a[x]
	}

	k := cnt
	for x := n - 1; x >= 1; x-- {
		for cnt > k && xmul(res[cnt-1], res[cnt], a[x]) < 0 {
			cnt--
		}
		cnt++
		res[cnt] = a[x]
	}

	sum := 0
	for x := 1; x <= cnt; x++ {
		sum += res[x].x*res[x%cnt+1].y - res[x].y*res[x%cnt+1].x
	}
	sum = abs(sum)
	sum2 := 0
	for x := 1; x <= cnt; x++ {
		sum2 += gcd(abs(res[x].x-res[x%cnt+1].x), abs(res[x].y-res[x%cnt+1].y))
	}
	fmt.Println((sum+sum2+2)/2 - n)
}

func xmul(a, b, c point) int {
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
