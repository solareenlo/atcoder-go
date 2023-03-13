package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		x, y int
	}

	var ask func(float64) float64
	ask = func(ang float64) float64 {
		var ymax float64
		fmt.Printf("? %.12f\n", ang)
		out.Flush()
		fmt.Scanf("%f", &ymax)
		return ymax
	}

	var deg [2000]float64
	V := make([]pair, 0)
	V = append(V, pair{0, 0})
	for i := 0; i <= 720; i++ {
		deg[i] = ask(float64(i) * 0.5)
		if i < 2 {
			continue
		}
		cy := float64(V[len(V)-1].y)
		cx := float64(V[len(V)-1].x)
		dc := float64(i) * 0.5 * math.Pi / 180.0
		dp := float64(i-1) * 0.5 * math.Pi / 180
		c1 := math.Cos(dc)
		s1 := math.Sin(dc)
		c2 := math.Cos(dp)
		s2 := math.Sin(dp)
		y1 := deg[i]
		y2 := deg[i-1]
		if y1 < cy*c1+cx*s1+0.00001 || y2 < cy*c2+cx*s2+0.00001 {
			continue
		}
		y := (y1*s2 - y2*s1) / (c1*s2 - c2*s1)
		x := (y1*c2 - y2*c1) / (c2*s1 - c1*s2)
		V = append(V, pair{int(math.Floor(x + 0.5)), int(math.Floor(y + 0.5))})
	}
	fmt.Printf("! %d\n", len(V)-1)
	for i := 0; i < len(V); i++ {
		fmt.Printf("! %d %d\n", V[len(V)-1-i].x, V[len(V)-1-i].y)
	}
}
