package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var dir [4][2]int = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var rest, v [10][10]int
	var totalCount int
	var h, w int
	dat := make([]string, 10)

	var findAns func(int, int, int, int) bool
	findAns = func(r, c, cnt, pdir int) bool {
		if cnt+1 == totalCount {
			return true
		}
		v[r][c] = 1
		for i := 0; i < 4; i++ {
			if i == pdir {
				continue
			}
			nr := r + dir[i][0]
			nc := c + dir[i][1]
			if nr < 0 || nc < 0 || nr >= h || nc >= w {
				continue
			}
			if dat[nr][nc] == '#' {
				continue
			}
			if rest[nr][nc] > cnt {
				continue
			}
			if v[nr][nc] == 0 {
				if findAns(nr, nc, cnt+1, i^1) {
					return true
				}
			}
		}
		v[r][c] = 0
		return false
	}

	for {
		var len int
		fmt.Fscan(in, &len, &h, &w)
		if len == 0 && h == 0 && w == 0 {
			break
		}
		hr, hc := 0, 0
		sr, sc := 0, 0
		totalCount = 0
		for i := 0; i < h; i++ {
			fmt.Fscan(in, &dat[i])
			for j := 0; j < w; j++ {
				rest[i][j] = 0
				v[i][j] = 0
				if dat[i][j] >= '0' && dat[i][j] <= '9' {
					rest[i][j] = len - int(dat[i][j]-'0')
					totalCount++
				} else if dat[i][j] == '.' {
					totalCount++
				}
				if dat[i][j] == '1' {
					hr = i
					hc = j
				}
				if dat[i][j] == '2' {
					sr = i
					sc = j
				}
			}
		}
		pdir := 0
		for i := 0; i < 4; i++ {
			if sr+dir[i][0] == hr && sc+dir[i][1] == hc {
				pdir = i
				break
			}
		}
		if findAns(hr, hc, 0, pdir^1) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
