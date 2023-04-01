package main

import (
	"fmt"
	"strconv"
)

func main() {
	type pair struct {
		x string
		y int
	}

	var s string
	fmt.Scan(&s)
	sz := len(s)
	var a int
	fmt.Scan(&a)
	p := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&p[i])
	}

	var dp [20][20][]pair
	for i := 0; i < a; i++ {
		str := strconv.Itoa(i + 1)
		dp[i][i] = append(dp[i][i], pair{str, p[i]})
	}

	for i := 1; i < a; i++ {
		for j := 0; j+i < a; j++ {
			for k := 1; k < i+1; k++ {
				for l := 0; l < len(dp[j][j+k-1]); l++ {
					for m := 0; m < len(dp[j+k][j+i]); m++ {
						Lw := dp[j][j+k-1][l].y
						Rw := dp[j+k][j+i][m].y
						if Lw == 0 && Rw == 0 {
							continue
						}
						if Lw != 0 && Rw != 0 {
							continue
						}
						if Lw != 0 {
							dp[j][i+j] = append(dp[j][i+j], pair{"[(" + dp[j][j+k-1][l].x + ")o][(" + dp[j+k][j+i][m].x + ")x]", Lw - 1})
						} else {
							dp[j][i+j] = append(dp[j][i+j], pair{"[(" + dp[j][j+k-1][l].x + ")x][(" + dp[j+k][j+i][m].x + ")o]", Rw - 1})
						}
					}
				}
			}
		}
	}

	for i := 0; i < len(dp[0][a-1]); i++ {
		S := dp[0][a-1][i].x
		if sz != len(S) {
			continue
		}
		ok := true
		for j := 0; j < len(S); j++ {
			if S[j] != s[j] && s[j] != '?' {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(S)
			return
		}
	}
}
