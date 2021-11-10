package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W, Ch, Cw, Dh, Dw int
	fmt.Fscan(in, &H, &W, &Ch, &Cw, &Dh, &Dw)

	s := make([]string, H)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}

	dist := make([][]int, H)
	for i := range dist {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = -1
		}
	}

	type tuple struct{ h, w, d int }
	deque := list.New()
	deque.PushBack(tuple{Ch - 1, Cw - 1, 0})
	for deque.Len() > 0 {
		t := deque.Front().Value.(tuple)
		h, w, d := t.h, t.w, t.d
		deque.Remove(deque.Front())
		if h < 0 || H <= h || w < 0 || W <= w {
			continue
		}
		if s[h][w] == '#' {
			continue
		}
		if dist[h][w] >= 0 {
			continue
		}
		dist[h][w] = d
		for dh := -2; dh <= 2; dh++ {
			for dw := -2; dw <= 2; dw++ {
				if dh*dh+dw*dw == 1 {
					deque.PushFront(tuple{h + dh, w + dw, d})
				} else if dh*dh+dw*dw > 1 {
					deque.PushBack(tuple{h + dh, w + dw, d + 1})
				}
			}
		}
	}

	fmt.Println(dist[Dh-1][Dw-1])
}
