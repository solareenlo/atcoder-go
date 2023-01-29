package main

import "fmt"

var n, m, nn, mm int
var a, b [11][11]int
var c, d [11]int
var ok bool

func main() {
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	fmt.Scan(&nn, &mm)
	for i := 1; i <= nn; i++ {
		for j := 1; j <= mm; j++ {
			fmt.Scan(&b[i][j])
		}
	}

	dfs(1, 1)
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func dfs(x, y int) {
	if x > nn {
		if y > mm {
			for i := 1; i <= nn; i++ {
				for j := 1; j <= mm; j++ {
					if b[i][j] != a[c[i]][d[j]] {
						return
					}
				}
			}
			ok = true

			return
		}
		for i := d[y-1] + 1; i <= m; i++ {
			d[y] = i
			dfs(x, y+1)
		}
	}
	for i := c[x-1] + 1; i <= n; i++ {
		c[x] = i
		dfs(x+1, y)
	}
}
