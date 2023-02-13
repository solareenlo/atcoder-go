package main

import "fmt"

var n, m int
var ma [30][30]byte
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}
var vis [30][30][30][4]bool

func main() {
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		var t string
		fmt.Scan(&t)
		for j := 0; j < len(t); j++ {
			ma[i][j] = t[j]
		}
	}
	if dfs(0, 0, 0, 3) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func dfs(x, y, v, d int) bool {
	x = (x + n) % n
	y = (y + m) % m
	if vis[x][y][v][d] {
		return false
	}
	vis[x][y][v][d] = true
	switch ma[x][y] {
	case '^':
		d = 0
	case 'v':
		d = 1
	case '<':
		d = 2
	case '>':
		d = 3
	case '_':
		if v == 0 {
			d = 3
		} else {
			d = 2
		}
	case '|':
		if v == 0 {
			d = 1
		} else {
			d = 0
		}
	case '?':
		for i := 0; i < 4; i++ {
			if dfs(x+dx[i], y+dy[i], v, i) {
				return true
			}
		}
	case '.':
	case '@':
		return true
	case '+':
		v = (v + 1) % 16
	case '-':
		v = (v + 15) % 16
	default:
		v = int(ma[x][y] - '0')
	}
	if dfs(x+dx[d], y+dy[d], v, d) {
		return true
	}
	return false
}
