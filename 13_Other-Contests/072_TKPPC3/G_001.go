package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const N = 200001
const ep = 1e-11

var a, b, m [N]float64
var cnt, q [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < N; i++ {
		m[i] = -1e9
	}
	for i := 0; i < n; i++ {
		var b int
		var a, r float64
		fmt.Fscan(in, &a, &b, &r)
		d := math.Log2(a) - math.Log2(float64(b))*r
		if d-m[b] > ep {
			m[b] = d
			cnt[b] = 1
		} else if math.Abs(d-m[b]) < ep {
			cnt[b]++
		}
	}
	s, t := 0, 0
	for i := 0; i < N; i++ {
		if m[i] != -1e9 {
			a[i] = math.Log2(float64(i))
			b[i] = m[i]
			for t-s > 1 && check(q[t-2], q[t-1], i) {
				t--
			}
			q[t] = i
			t++
		}
	}
	var k int
	fmt.Fscan(in, &k)
	ans := 0
	for i := 0; i < k; i++ {
		var y float64
		fmt.Fscan(in, &y)
		for t-s > 1 && (a[q[s+1]]*y+b[q[s+1]])-(a[q[s]]*y+b[q[s]]) > ep {
			s++
		}
		for t-s > 1 && cnt[q[s]] == 0 && math.Abs((a[q[s+1]]*y+b[q[s+1]])-(a[q[s]]*y+b[q[s]])) < ep {
			s++
		}
		if cnt[q[s]] != 0 {
			ans++
			cnt[q[s]]--
		}
	}
	fmt.Println(ans)
}

func check(i, j, k int) bool {
	return (a[j]-a[i])*(b[k]-b[j])-(a[k]-a[j])*(b[j]-b[i]) > ep
}
