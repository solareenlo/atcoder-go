package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	que := list.New()
	di := make([]int, n)
	for i := 0; i < n; i++ {
		di[i] = 1 << 60
	}
	di[1] = 1
	que.PushFront(1)
	for que.Len() > 0 {
		x := que.Front().Value.(int)
		que.Remove(que.Front())
		if x == 0 {
			fmt.Println(di[0])
			return
		}
		nx := x + 1
		if nx == n {
			nx = 0
		}
		if di[nx] > di[x]+1 {
			di[nx] = di[x] + 1
			que.PushBack(nx)
		}
		nx = (x * 10) % n
		if di[nx] > di[x] {
			di[nx] = di[x]
			que.PushFront(nx)
		}
	}
}
