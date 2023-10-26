package main

import (
	"fmt"
	"math"
)

type pair struct {
	x, y float64
}

var PX, PY, PR float64

func angle(ax, ay, bx, by float64) float64 {
	ret := math.Atan2(by, bx) - math.Atan2(ay, ax)
	ret /= 8 * math.Atan(1)
	if ret < 0 {
		ret += 1
	}
	return ret
}

func hoge(X, Y float64, dir int) pair {
	r := math.Hypot(X-PX, Y-PY)
	deg := math.Atan2(PY-Y, PX-X) + float64(dir)*math.Asin(PR/r)
	dx := math.Cos(deg)
	dy := math.Sin(deg)
	di := 2.0
	for i := 0; i < 100; i++ {
		nx := X + di*dx
		ny := Y + di*dy
		if ny*ny+nx*nx <= 1 {
			X = nx
			Y = ny
		}
		di /= 2
	}
	return pair{X, Y}
}

func main() {
	fmt.Scan(&PX, &PY, &PR)
	var SX, SY float64
	var K int
	fmt.Scan(&SX, &SY, &K)

	a := hoge(SX, SY, 1)
	var LX, LY [11]float64
	LX[0] = a.x
	LY[0] = a.y
	a = hoge(SX, SY, -1)
	var RX, RY [11]float64
	RX[0] = a.x
	RY[0] = a.y

	ret := angle(LX[0], LY[0], RX[0], RY[0])
	fmt.Println(math.Min(1.0, ret))
	for i := 0; i < K; i++ {
		a = hoge(LX[i], LY[i], 1)
		LX[i+1] = a.x
		LY[i+1] = a.y
		a = hoge(RX[i], RY[i], -1)
		RX[i+1] = a.x
		RY[i+1] = a.y
		ret += angle(LX[i+1], LY[i+1], LX[i], LY[i])
		ret += angle(RX[i], RY[i], RX[i+1], RY[i+1])
		fmt.Println(math.Min(1.0, ret))
	}
}
