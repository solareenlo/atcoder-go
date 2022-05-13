package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	S := make([]string, N)
	for i := range S {
		fmt.Scan(&S[i])
	}

	ans := 0
	cnt := make([]int, 26)
	for i := 0; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			if i>>j&1 != 0 {
				for _, c := range S[j] {
					cnt[c-'a']++
				}
			}
		}
		now := 0
		for j := 0; j < 26; j++ {
			if cnt[j] == K {
				now++
			}
			cnt[j] = 0
		}
		if ans < now {
			ans = now
		}
	}
	fmt.Println(ans)
}
