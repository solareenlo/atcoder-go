package main

import (
	"bufio"
	"fmt"
	"os"
)

var ans [500000]int
var par [101]int

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Scan(&t)
	if t == 1 {
		fmt.Fprintln(out, "300000")
		fmt.Fprintf(out, "2 2 1 2 1 2 1 1 2 3 1 1")
		for i := 12; i < 300000; i++ {
			fmt.Fprintf(out, " 1")
		}
		fmt.Fprintln(out)
	} else {
		var i, j int
		for i = 0; i <= 100; i++ {
			par[i] = i
		}
		ans[0] = 100
		ans[1] = 100
		for i, j = 2, 99; i <= 100; i++ {
			if i%12 == 1 && j+j-1 > 100 {
				ans[i] = j + j - 1
			} else if i%12 == 4 && ans[i-3] > 100 {
				ans[i] = ans[i-3] - 1
			} else {
				ans[i] = j
				j--
			}
		}
		for i = 2; i <= 100; i += 3 {
			unite(ans[i], ans[i+1])
		}
		for i, j = 101, 99; j >= 1; j-- {
			if find(100) == find(j) {
				continue
			}
			unite(100, j)
			ans[i] = 100
			i++
			ans[i] = j
			i++
			ans[i] = 100
			i++
		}
		for i = 101; i <= 298; i++ {
			if i%12 == 1 {
				ans[i] = ans[i+1] + ans[i+2]
			}
			if i%12 == 4 {
				ans[i] = ans[i-3]
			}
		}
		ans[i] = 100
		i++
		ans[i] = 1
		i++
		for j = 0; j < 25; j++ {
			ans[i] = 200
			i++
			ans[i] = 100
			i++
			ans[i] = 100
			i++
			ans[i] = 200
			i++
		}
		for ; i < 500000; i++ {
			ans[i] = 1e9
		}
		fmt.Fprintln(out, "500000")
		for i = 0; i < 500000; i++ {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, "%d", ans[i])
		}
		fmt.Fprintln(out)
	}
}

func find(x int) int {
	if par[x] == x {
		return x
	}
	par[x] = find(par[x])
	return par[x]
}

func unite(x, y int) {
	x = find(x)
	y = find(y)
	par[x] = y
}
