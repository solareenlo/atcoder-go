package main

import (
	"fmt"
)

var H, W int

func issame(x, y, l, u int) bool {
	if x%2+y%2 == 2 {
		return true
	}
	if x%2+y%2 == 1 {
		return false
	}
	for l <= u {
		fmt.Printf("? %d %d %d %d\n", l+1, x/2+1, u+1, y/2+1)
		var s string
		fmt.Scan(&s)
		if s[0] == 'n' {
			return false
		}
		l++
		u--
	}
	return true
}

func main() {
	fmt.Scan(&H, &W)
	var ans int64 = 0
	N := 2*W - 1
	for a := 0; a < H; a++ {
		for b := a; b < H; b++ {
			R := make([]int, 200)
			i := 0
			j := 0
			for i < N {
				for i-j >= 0 && i+j < N && issame(i-j, i+j, a, b) {
					j++
				}
				R[i] = j
				if R[i] == 0 {
					i++
					j = 0
					continue
				}
				k := 1
				for i-k >= 0 && i+k < N && k+R[i-k] < j {
					R[i+k] = R[i-k]
					k++
				}
				i += k
				j -= k
			}
			for i := 0; i < N; i++ {
				if i%2 == 1 {
					ans += int64(R[i] / 2)
				} else {
					ans += int64((R[i] + 1) / 2)
				}
			}
		}
	}
	fmt.Printf("! %d\n", ans)
}
