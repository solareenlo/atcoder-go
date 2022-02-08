package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const N = 210

var (
	n   int
	vis = make([]bool, N)
	bel = make([]int, N)
	G   = [N][N]bool{}
)

func find(i int) bool {
	for j := 1; j <= n; j++ {
		if !vis[j] && G[i][j] {
			vis[j] = true
			if bel[j] == 0 || find(bel[j]) {
				bel[j] = i
				return true
			}
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	m := map[int]int{}
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		m[x] ^= 1
		m[x+1] ^= 1
	}

	a := make([]int, N)
	b := make([]int, N)
	for k, v := range m {
		if v != 0 {
			if k&1 != 0 {
				a[0]++
				a[a[0]] = k
			} else {
				b[0]++
				b[b[0]] = k
			}
		}
	}

	ans := 0
	for i := 1; i <= a[0]; i++ {
		for j := 1; j <= b[0]; j++ {
			if isPrime(abs(a[i] - b[j])) {
				G[i][j] = true
			}
		}
	}

	for i := 1; i <= a[0]; i++ {
		for j := range vis {
			vis[j] = false
		}
		if find(i) {
			ans++
		}
	}

	fmt.Println(ans + (a[0]-ans)/2*2 + (b[0]-ans)/2*2 + (a[0]-ans)%2*3)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqr := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqr; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
