package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const INF int = (1 << 29)

var N, Q int
var a [4010]int
var qs [200010]int
var qt [200010]int
var ans [200010]int
var dist [4010][4010]int

func func2(x int) {
	var y int
	for i := 0; i < N; i++ {
		dist[x][i] = int(math.Abs(float64(x-i))) * a[x]
	}
	for y = x - 1; y >= 0; y-- {
		if a[y] < a[x] {
			break
		}
	}
	if y >= 0 {
		for i := 0; i < N; i++ {
			dist[x][i] = int(math.Min(float64(dist[x][i]), float64(dist[x][y]+dist[y][i])))
		}
	}
	for y = x + 1; y < N; y++ {
		if a[y] < a[x] {
			break
		}
	}
	if y < N {
		for i := 0; i < N; i++ {
			dist[x][i] = int(math.Min(float64(dist[x][i]), float64(dist[x][y]+dist[y][i])))
		}
	}
}

func func1() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			dist[i][j] = INF
		}
	}
	var v []struct {
		first  int
		second int
	}
	for i := 0; i < N; i++ {
		v = append(v, struct {
			first  int
			second int
		}{a[i], i})
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].first < v[j].first
	})
	for i := 0; i < N; i++ {
		x := v[i].second
		func2(x)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &N, &Q)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &qs[i], &qt[i])
		qs[i]--
		qt[i]--
	}
	func1()
	for i := 0; i < Q; i++ {
		ans[i] = dist[qt[i]][qs[i]]
	}
	for i := 0; i < Q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
