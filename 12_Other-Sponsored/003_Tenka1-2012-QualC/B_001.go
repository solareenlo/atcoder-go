package main

import "fmt"

func main() {

	const SHDC = "SHDC"

	var poi [4]string
	var count [4]int

	var S string
	fmt.Scan(&S)
	N := len(S)
	for j := 0; j < N; j++ {
		for i := 0; i < 4; i++ {
			if ('2' <= S[j+1] && S[j+1] <= '9') || (S[j] != SHDC[i]) {
				poi[i] += string(S[j])
				poi[i] += string(S[j+1])
				if S[j+1] == '1' {
					poi[i] += string(S[j+2])
				}
				continue
			}
			count[i]++
			if count[i] == 5 {
				if len(poi[i]) == 0 {
					fmt.Println(0)
				} else {
					fmt.Println(poi[i])
				}
				return
			}
		}
		if S[j+1] == '1' {
			j += 2
		} else {
			j++
		}
	}
	fmt.Println(-1)
}
