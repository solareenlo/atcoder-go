package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

type Pair struct {
	a, b int
}

var p [3005]Pair
var perm []int

func main() {
	in := bufio.NewReader(os.Stdin)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	score := -1.0

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &p[i].a, &p[i].b)
	}
	perm = make([]int, N)
	for i := 0; i < N; i++ {
		perm[i] = i
	}
	ans := make([]int, N)
	for x := 0; x < 100; x++ {
		rand.Shuffle(len(perm), func(i, j int) {
			perm[i], perm[j] = perm[j], perm[i]
		})
		sum := 0.0
		for i := 0; i < N; i += 3 {
			sum += getArea(i)
		}
		changed := 0
		for y := 0; y < 5000; y++ {
			a := r.Int() % N
			b := r.Int() % N
			if a/3 == b/3 {
				continue
			}
			A := getArea(a / 3 * 3)
			B := getArea(b / 3 * 3)
			perm[a], perm[b] = perm[b], perm[a]
			nA := getArea(a / 3 * 3)
			nB := getArea(b / 3 * 3)
			if A+B < nA+nB {
				sum += nA + nB - A - B
				changed++
			} else {
				perm[a], perm[b] = perm[b], perm[a]
			}
		}
		if score < sum {
			score = sum
			for i := 0; i < N; i++ {
				ans[i] = perm[i]
			}
		}
	}
	for i := 0; i < N; i += 3 {
		fmt.Printf("%d %d %d\n", ans[i], ans[i+1], ans[i+2])
	}
}

func getArea(i int) float64 {
	a := getLen(perm[i], perm[i+1])
	b := getLen(perm[i+1], perm[i+2])
	c := getLen(perm[i], perm[i+2])
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func getLen(a, b int) float64 {
	return math.Sqrt(sq(p[a].a-p[b].a) + sq(p[a].b-p[b].b))
}

func sq(a int) float64 { return float64(a * a) }
