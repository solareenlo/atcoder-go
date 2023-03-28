package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const nax = 57
const dlu = 10

var n int = 50
var wyn, mod int
var pot, jeszkol [nax][nax]int
var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var out = bufio.NewWriter(os.Stdout)

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z, &mod)
	trz := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			pot[i][j] = trz
			trz = (trz * 2) % mod
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < nax; j++ {
			if dlu*j < n {
				jeszkol[i][j] = pyt(dlu*j, min(dlu*j+dlu-1, n-1), i, i)
			}
		}
	}
	szuk(0, n-1, 0, n-1, 250)
	fmt.Printf("! %d\n", wyn%mod)
	out.Flush()
}

func pyt(a, b, x, y int) int {
	fmt.Printf("? %d %d %d %d\n", a, x, b, y)
	out.Flush()
	var v int
	fmt.Scan(&v)
	return v
}

func szuk(a, b, x, y, w int) {
	if w == 0 {
		return
	}
	if (a/dlu) == (b/dlu) && jeszkol[x][a/dlu] == 0 {
		szuk(a, b, x+1, y, w)
		return
	}
	if (a/dlu) == (b/dlu) && jeszkol[y][a/dlu] == 0 {
		szuk(a, b, x, y-1, w)
		return
	}
	if a == b {
		ile := 0
		for i := x; i <= y; i++ {
			if jeszkol[i][a/dlu] > 0 {
				ile++
			}
		}
		if ile == w {
			for j := x; j <= y; j++ {
				if jeszkol[j][a/dlu] != 0 {
					wyn += pot[a][j]
					jeszkol[j][a/dlu]--
				}
			}
			return
		}
	}
	if w == (b-a+1)*(y-x+1) {
		for i := a; i <= b; i++ {
			for j := x; j <= y; j++ {
				wyn += pot[i][j]
				jeszkol[j][i/dlu]--
			}
		}
		return
	}
	if a != b {
		s := (a + b) >> 1
		p := pyt(a, s, x, y)
		if (r.Int31() & 1) != 0 {
			szuk(a, s, x, y, p)
			szuk(s+1, b, x, y, w-p)
		} else {
			szuk(s+1, b, x, y, w-p)
			szuk(a, s, x, y, p)
		}
	} else {
		s := (x + y) >> 1
		p := pyt(a, b, x, s)
		if (r.Int31() & 1) != 0 {
			szuk(a, b, x, s, p)
			szuk(a, b, s+1, y, w-p)
		} else {
			szuk(a, b, s+1, y, w-p)
			szuk(a, b, x, s, p)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
