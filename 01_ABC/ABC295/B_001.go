package main

import (
	"fmt"
	"strings"
)

func main() {
	var s [20][]string
	var a, b int
	fmt.Scan(&a, &b)
	for i := 0; i < a; i++ {
		var tmp string
		fmt.Scan(&tmp)
		s[i] = strings.Split(tmp, "")
	}
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if s[i][j] > "0" && s[i][j] <= "9" {
				c := int(s[i][j][0] - '0')
				for k := 0; k < a; k++ {
					for l := 0; l < b; l++ {
						if abs(i-k)+abs(j-l) <= c && !(s[k][l][0] > '0' && s[k][l][0] <= '9') || k == i && j == l {
							s[k][l] = "."
						}
					}
				}
			}
		}
	}
	for i := 0; i < a; i++ {
		fmt.Println(strings.Join(s[i], ""))
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
