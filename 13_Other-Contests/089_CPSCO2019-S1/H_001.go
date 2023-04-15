package main

import (
	"bufio"
	"fmt"
	"os"
)

var p [110000]int
var X, ret int
var val [310000]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)
	X = b
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &p[i])
	}
	for i := 0; i < a; i++ {
		if p[i]*3 == b {
			ret++
		}
	}
	calc(0, a)
	fmt.Println(ret)
}

func calc(a, b int) {
	if b-a < 2 {
		return
	}
	M := (a + b) / 2
	calc(a, M)
	calc(M, b)
	ks := 0
	h := -1
	at := M - 1
	for i := M; i < b; i++ {
		if h <= p[i] {
			if h != p[i] {
				ks = 1
			} else {
				ks++
			}
			h = p[i]
		}
		for at >= a && p[at] <= h {
			val[p[at]]++
			at--
		}
		if h >= p[i] {
			if X-h-p[i] >= 0 {
				ret += ks * val[X-h-p[i]]
			}
		}
	}
	for i := at + 1; i < M; i++ {
		val[p[i]]--
	}
	h = -1
	ks = 0
	at = M
	for i := M - 1; i >= a; i-- {
		if h <= p[i] {
			if h != p[i] {
				ks = 1
			} else {
				ks++
			}
			h = p[i]
		}
		for at < b && p[at] <= h {
			val[p[at]]++
			at++
		}
		if h >= p[i] {
			if X-h-p[i] >= 0 {
				ret += ks * val[X-h-p[i]]
			}
		}
	}
	for i := M; i < at; i++ {
		val[p[i]]--
	}
}
