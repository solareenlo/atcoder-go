package main

import (
	"container/list"
	"fmt"
)

type t struct {
	x, y, color int
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([]string, h)
	white := 0
	for i := range s {
		fmt.Scan(&s[i])
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				white++
			}
		}
	}

	deque := list.New()
	if s[0][0] == '.' {
		deque.PushFront(t{0, 0, 1})
	}
	black := 0
	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	for deque.Len() > 0 {
		f := deque.Remove(deque.Front()).(t)
		if f.x == h-1 && f.y == w-1 {
			black = f.color
			break
		}
		for i := 0; i < 4; i++ {
			nx, ny := f.x+dx[i], f.y+dy[i]
			if 0 <= nx && nx < h && 0 <= ny && ny < w && s[nx][ny] == '.' {
				deque.PushBack(t{nx, ny, f.color + 1})
				s[nx] = s[nx][:ny] + "#" + s[nx][ny+1:]
			}
		}
	}
	if black != 0 {
		fmt.Println(white - black)
	} else {
		fmt.Println(-1)
	}
}
