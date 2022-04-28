package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	ss := make([]string, 18)
	for i := range ss {
		fmt.Fscan(in, &ss[i])
	}
	ss = append(ss, "000000")
	ss = append(ss, "000000")

	dp := make([]int, 1<<12)
	dp[0] = 1
	mask := len(dp) - 1
	for _, t := range ss {
		for i := 0; i < len(t); i++ {
			nx := make([]int, 1<<12)
			for s := 0; s != len(dp); s++ {
				cnt := [2]int{}
				if i != 0 {
					cnt[s>>6&1]++
				} else {
					cnt[0]++
				}
				if i < 5 {
					cnt[s>>4&1]++
				} else {
					cnt[0]++
				}
				cnt[s>>11&1]++
				up := (s >> 5) & 1
				cnt[up]++
				for f := 0; f < 2; f++ {
					if int(t[i]) != ('0' ^ f ^ 1) {
						cnt[f]++
						if cnt[up] > cnt[1^up] {
							nx[mask&(s<<1|f)] += dp[s]
						}
						cnt[f]--
					}
				}
			}
			dp = nx
		}
	}
	fmt.Println(dp[0])
}
