package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var n int
var a [5005]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	T := 20
	rand.Seed(time.Now().UnixNano())
	for T > 0 {
		x := rand.Intn(n)*rand.Intn(n)%n + 1
		y := rand.Intn(n)*rand.Intn(n)%n + 1
		d := abs(a[x] - a[y])
		for i := 1; i*i <= d; i++ {
			if d%i == 0 {
				check(i, a[x]%i)
				if i*i != d {
					check(d/i, a[x]%(d/i))
				}
			}
		}
		T--
	}
	fmt.Println(-1)
}

func check(x, y int) {
	if x < 3 {
		return
	}
	r := 0
	for i := 1; i <= n; i++ {
		if a[i]%x == y {
			r++
		} else {
			r--
		}
	}
	if r > 0 {
		fmt.Println(x)
		os.Exit(0)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
