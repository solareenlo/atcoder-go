package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(100001))

	type tuple struct {
		x, y, z int
	}

	var x float64
	fmt.Scan(&x)

	x2 := x * x
	x3 := x * x * x

	memo := make(map[int]tuple)
	for {
		a := r.Intn(1000)
		b := r.Intn(10000)
		c := r.Intn(100000)
		fx := float64(a)*x3 + float64(b)*x2 + float64(c)*x
		d := math.Floor(fx)
		key := int(2.0 * math.Pow(10, 10) * (fx - d))
		if _, ok := memo[key]; ok {
			t := memo[key]
			a1 := t.x
			b1 := t.y
			c1 := t.z
			if a == a1 && b == b1 && c == c1 {
				continue
			}
			a = a - a1
			b = b - b1
			c = c - c1
			fx = float64(a)*x3 + float64(b)*x2 + float64(c)*x
			d = math.Round(fx)
			if math.Abs(fx-d) < 9e-11 {
				fmt.Println(a, b, c, -d)
				return
			}
		} else {
			memo[key] = tuple{a, b, c}
		}
	}
}
