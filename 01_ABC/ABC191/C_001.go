package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([]string, h)
	for i := range s {
		fmt.Scan(&s[i])
	}

	cnt := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			if s[i][j]^s[i+1][j]^s[i][j+1]^s[i+1][j+1] != 0 {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
