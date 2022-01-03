package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	c = make([][]byte, 1002)
)

func dfs(x, y int) int {
	ret := 0
	if c[x][y] != 'o' {
		return 0
	}
	c[x][y] = '.'
	ret++
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ret += dfs(x+i-1, y+j-1)
		}
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(in, &tmp)
		tmp = " " + tmp
		c[i+1] = []byte(tmp)
	}

	ans := [3]int{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if c[i+1][j+1] == 'o' {
				t := dfs(i+1, j+1)
				for i := 1; i*i < t; i++ {
					if t%(i*i) == 0 {
						u := t / (i * i)
						if u == 12 {
							ans[0]++
							break
						} else if u == 16 {
							ans[1]++
							break
						} else if u == 11 {
							ans[2]++
							break
						}
					}
				}
			}
		}
	}

	fmt.Println(ans[0], ans[1], ans[2])
}
