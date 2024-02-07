package main

import (
	"fmt"
	"os"
	"strings"
)

const N = 100

var n int
var ans [N][N]string
var tx, ty [N][N]bool
var fx, fy [N]bool
var a, b string

func dfs(x, y, tot int) {
	if y > n {
		x++
		y = 1
	}
	if x > n {
		if tot == 3*n {
			fmt.Println("Yes")
			for i := 1; i <= n; i++ {
				fmt.Println(strings.Join(ans[i][1:], ""))
			}
			os.Exit(0)
		}
		return
	}
	for i := 0; i < 3; i++ {
		if !tx[x][i] && !ty[y][i] {
			flag1 := fx[x]
			flag2 := fy[y]
			if !fx[x] && i != int(a[x]-'A') {
				continue
			}
			if !fy[y] && i != int(b[y]-'A') {
				continue
			}
			fx[x] = true
			fy[y] = true
			tx[x][i] = true
			ty[y][i] = true
			ans[x][y] = string('A' + i)
			dfs(x, y+1, tot+1)
			fx[x] = flag1
			fy[y] = flag2
			tx[x][i] = false
			ty[y][i] = false
		}
	}
	ans[x][y] = "."
	dfs(x, y+1, tot)
}

func main() {
	fmt.Scan(&n, &a, &b)
	a = " " + a
	b = " " + b
	dfs(1, 1, 0)
	fmt.Println("No")
}
