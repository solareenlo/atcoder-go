package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a [200002]int
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[x]++
	}

	var b [200002]int
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(in, &x)
		b[x]++
	}

	v := make([]int, 0)
	w := make([]int, 0)
	z := make([]int, 0)
	for i := 1; i <= n; i++ {
		if a[i]+b[i] <= 1 {
			fmt.Println("No")
			return
		}
		if a[i] > 0 && b[i] > 0 {
			z = append(z, i)
		} else if a[i] > 0 {
			v = append(v, i)
		} else {
			w = append(w, i)
		}
	}

	tmp1 := make([]int, 0)
	tmp1 = append(tmp1, -1)
	tmp2 := make([]int, 0)
	for i := 0; i < len(z); i++ {
		tmp1 = append(tmp1, z[i])
		tmp2 = append(tmp2, z[i])
		a[z[i]]--
		b[z[i]]--
	}

	for i := 0; i < min(len(v), len(w)); i++ {
		tmp1 = append(tmp1, v[i])
		tmp1 = append(tmp1, v[i])
		tmp2 = append(tmp2, w[i])
		tmp2 = append(tmp2, w[i])
		a[v[i]] -= 2
		b[w[i]] -= 2
	}

	tmp2 = append(tmp2, tmp2[0])

	ans := make([][]pair, 2)
	d := 0
	for i := 1; i < len(tmp1); i++ {
		ans[d] = append(ans[d], pair{tmp1[i], tmp2[i]})
		if i >= len(z) {
			d = 1 - d
		}
	}

	tmp1 = make([]int, 0)
	tmp2 = make([]int, 0)
	for i := 1; i <= n; i++ {
		for j := 0; j < a[i]; j++ {
			tmp1 = append(tmp1, i)
		}
		for j := 0; j < b[i]; j++ {
			tmp2 = append(tmp2, i)
		}
	}

	for i := 0; i < len(tmp1); i++ {
		ans[d] = append(ans[d], pair{tmp1[i], tmp2[i]})
		d = 1 - d
	}

	fmt.Println("Yes")
	for i := 0; i < len(ans[0]); i++ {
		fmt.Printf("%d > %d\n", ans[0][i].x, ans[0][i].y)
	}
	for i := 0; i < len(ans[1]); i++ {
		fmt.Printf("%d < %d\n", ans[1][i].x, ans[1][i].y)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
