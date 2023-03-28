package main

import (
	"fmt"
)

var (
	S   string
	idx int
)

func expression(low, high int) (int, int) {
	last := 1 << 30
	if S[idx] >= '0' && S[idx] <= '9' {
		left := int(S[idx] - '0')
		right := 0
		if left != 0 {
			right = left*10 + 9
		}
		if right <= low || high <= left {
			last = idx
		}
		idx++
		if S[idx] >= '0' && S[idx] <= '9' {
			left = left*10 + int(S[idx]-'0')
			last = min(last, idx)
			idx++
		} else {
			last = min(last, idx)
		}
		return left, last
	} else {
		ismin := S[idx] == '_'
		idx += 2
		latteLeft, latteLast := expression(low, high)
		if ismin {
			high = min(high, latteLeft)
		} else {
			low = max(low, latteLeft)
		}
		if low >= high {
			last = latteLast
		}
		idx++
		maltaLeft, maltaLast := expression(low, high)
		last = min(last, maltaLast)
		idx++
		if ismin {
			return min(latteLeft, maltaLeft), last
		} else {
			return max(latteLeft, maltaLeft), last
		}
	}
}

func main() {
	var Q int
	fmt.Scan(&Q)
	for i := 0; i < Q; i++ {
		fmt.Scan(&S)
		idx = 0
		left, right := expression(0, 99)
		fmt.Println(left, right+1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
