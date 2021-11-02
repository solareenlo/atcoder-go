package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var H, W, K int
	fmt.Scan(&H, &W, &K)

	S := make([]string, H)
	for i := range S {
		fmt.Scan(&S[i])
	}

	res := int(1e9)
	for bit := 0; bit < 1<<H; bit++ {
		t := bits.OnesCount(uint(bit))
		cnt := [11]int{}
		for i := 0; i < W; i++ {
			k := 0
			tmp := [11]int{}
			for j := 0; j < H; j++ {
				k += bit >> j & 1
				tmp[k] += int(S[j][i] - 48)
				if tmp[k] > K {
					goto L1
				}
				if cnt[k]+tmp[k] > K {
					t++
					for z := 0; z < 10; z++ {
						cnt[z] = 0
					}
				}
			}
			for j := 0; j < 10; j++ {
				cnt[j] += tmp[j]
			}
		}
		if res > t {
			res = t
		}
	L1:
	}
	fmt.Println(res)
}
