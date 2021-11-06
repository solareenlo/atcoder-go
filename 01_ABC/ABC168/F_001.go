package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type line struct{ a, b, c int }

// 上下左右に１マスずつ移動するための配列
// dx と dy は1つ要素がズレてるだけ．
var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, -1, 0, 1}
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	mapX := map[int]int{}
	mapY := map[int]int{}
	mapX[0] = 0
	mapY[0] = 0

	// 座標読み込み
	lineH := make([]line, 0)
	lineV := make([]line, 0)
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		mapX[a] = 0
		mapX[b] = 0
		mapY[c] = 0
		lineH = append(lineH, line{a, b, c})
	}
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &c, &a, &b)
		mapY[a] = 0
		mapY[b] = 0
		mapX[c] = 0
		lineV = append(lineV, line{a, b, c})
	}

	xs := make([]int, 0)
	ys := make([]int, 0) // 元の座標を求める用の座標
	// 圧縮用のマップ作成
	// 画像圧縮とは x が左から何本目ですか？が分かれば良い．
	keysX := make([]int, 0, len(mapX))
	for k := range mapX {
		keysX = append(keysX, k)
	}
	sort.Ints(keysX)
	for _, k := range keysX { // xs[mapX] とすると元のマップでの座標が分かる
		mapX[k] = len(xs)
		xs = append(xs, k)
	}

	keysY := make([]int, 0, len(mapY))
	for k := range mapY {
		keysY = append(keysY, k)
	}
	sort.Ints(keysY)
	for _, k := range keysY {
		mapY[k] = len(ys)
		ys = append(ys, k)
	}

	// 扱うグリッド．
	// 壁と普通の升目で2倍にする．
	h := len(xs) * 2
	w := len(ys) * 2

	d := make([][]int, h)
	for i := range d {
		d[i] = make([]int, w) // bfs 用の配列
	}

	// それぞれの線を引いていく
	// 横線を引く
	for i := 0; i < n; i++ {
		a := mapX[lineH[i].a] * 2
		b := mapX[lineH[i].b] * 2
		c := mapY[lineH[i].c] * 2
		for x := a; x <= b; x++ {
			d[x][c] = -1 // -1 は壁
		}
	}
	// 縦線を引く
	for i := 0; i < m; i++ {
		a := mapY[lineV[i].a] * 2
		b := mapY[lineV[i].b] * 2
		c := mapX[lineV[i].c] * 2
		for y := a; y <= b; y++ {
			d[c][y] = -1 // -1 は壁
		}
	}

	// bfs 開始
	q := &Heap{}
	sx := mapX[0] * 2
	sy := mapY[0] * 2
	d[sx][sy] = 1 // たどり着けのを1とする
	heap.Push(q, pair{sx, sy})
	for q.Len() > 0 {
		x := (*q)[0].x
		y := (*q)[0].y
		heap.Pop(q)
		// 上下左右を見る
		for v := 0; v < 4; v++ {
			nextX := x + dx[v]
			nextY := y + dy[v]
			if nextX < 0 || nextX >= h {
				continue
			}
			if nextY < 0 || nextY >= w {
				continue
			}
			if d[nextX][nextY] != 0 {
				continue
			}
			d[nextX][nextY] = 1
			heap.Push(q, pair{nextX, nextY})
		}
	}

	res := 0
	// 全部のマスを確認する
	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {
			if d[x][y] != 1 {
				continue // そのマスが到達不可能であれば無視する
			}
			if x == 0 || x == h-1 || y == 0 || y == w-1 { // 端に到達すれば INF
				fmt.Println("INF")
				return
			}
			// x か y が偶数なら壁なので飛ばす
			if x%2 == 0 || y%2 == 0 {
				continue // 今いるところが壁の行なら飛ばす
			}
			ex := xs[x/2+1] - xs[x/2] // xs[x / 2] が今いる場所の左下の元の座標
			ey := ys[y/2+1] - ys[y/2]
			res += ex * ey
		}
	}
	fmt.Println(res)
}

type pair struct{ x, y int }
type Heap []pair

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
