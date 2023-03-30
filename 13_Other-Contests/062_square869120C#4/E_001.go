package main

import (
	"fmt"
)

const P = 1000000007

func fmo(x *int) {
	*x += (*x >> 31) & P
}

func main() {
	var n, k, x int
	fmt.Scanf("%d%d%d", &n, &k, &x)
	f := make([][]int, x+x+1)
	for i := 0; i <= x+x; i++ {
		f[i] = make([]int, 10005)
	}
	f[0][1] = 1
	for i := 1; i <= x+x; i++ {
		o := i + (i & 1)
		for j, s := 1, 0; j <= 1e4; j++ {
			s += f[i-1][j-1] - P
			fmo(&s)
			if j-o-1 >= 0 {
				s -= f[i-1][j-o-1]
				fmo(&s)
			}
			f[i][j] = s
			if i+1 >= j {
				f[i][j] += 1 - P
				fmo(&f[i][j])
			}
		}
	}
	res := 1
	var a, b int
	fmt.Scanf("%d", &a)
	for i := 1; i < k; i++ {
		fmt.Scanf("%d", &b)
		res = (res * f[x+x-i-i+1][b-a]) % P
		a = b
	}
	fmt.Printf("%d\n", res)
}
