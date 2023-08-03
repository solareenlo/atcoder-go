package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, vcnt int
var v, match, matched [202]int
var graph [202][202]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	dat := make([]string, 0)
	for i := 0; i < n; i++ {
		var buff string
		fmt.Fscan(in, &buff)
		dat = append(dat, buff)
	}
	maxpos := len(dat[0])
	for i := 1; i < n; i++ {
		maxpos = min(maxpos, len(dat[i]))
	}
	for c := 'a'; c <= 'z'; c++ {
		for i := 0; i < n; i++ {
			match[i] = -1
			matched[i] = -1
		}
		for i := 0; i < n; i++ {
			for j := 0; j < min(n, len(dat[i])); j++ {
				graph[j][i] = (rune(dat[i][j]) == c)
			}
		}
		curcnt := 0
		for i := 0; i < n; i++ {
			vcnt++
			curcnt += findway(i)
		}
		for base := 0; base < maxpos; base++ {
			if curcnt == n {
				fmt.Println("YES")
				return
			}
			if base+1 == maxpos {
				break
			}
			kesu_node := base % n
			new_column := base + n
			for i := 0; i < n; i++ {
				graph[kesu_node][i] = (len(dat[i]) > new_column && rune(dat[i][new_column]) == c)
			}
			if match[kesu_node] != -1 {
				aite := match[kesu_node]
				match[kesu_node] = -1
				matched[aite] = -1
				curcnt--
			}
			for i := 0; i < n; i++ {
				if match[i] != -1 {
					continue
				}
				vcnt++
				if findway(i) != 0 {
					curcnt++
				} else {
					break
				}
			}
		}
	}
	fmt.Println("NO")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findway(nod int) int {
	if v[nod] == vcnt {
		return 0
	}
	v[nod] = vcnt
	for i := 0; i < n; i++ {
		if matched[i] == -1 && graph[nod][i] {
			matched[i] = nod
			match[nod] = i
			return 1
		} else if graph[nod][i] {
			if findway(matched[i]) != 0 {
				match[nod] = i
				matched[i] = nod
				return 1
			}
		}
	}
	return 0
}
