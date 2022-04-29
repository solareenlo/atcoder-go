package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)
	S := make([]string, 10)
	for i := 0; i < H; i++ {
		fmt.Scan(&S[i])
	}
	T := make([]string, 10)
	for i := 0; i < H; i++ {
		fmt.Scan(&T[i])
	}

	ok := false
	pH := H
	pW := W
	U := make([]string, 10)
	for r := 0; r < 4; r++ {
		for a := -10; a <= 10; a++ {
			for b := -10; b <= 10; b++ {
				now := true
				for i := 0; i < H; i++ {
					for j := 0; j < W; j++ {
						if T[i][j] == '#' {
							x := a + i
							y := b + j
							if x < 0 || y < 0 || x >= pH || y >= pW {
								now = false
							} else if S[x][y] == '#' {
								now = false
							}
						}
					}
				}
				if now {
					ok = true
				}
			}
		}
		for i := 0; i < W; i++ {
			U[i] = ""
			for j := 0; j < H; j++ {
				U[i] = U[i] + string(T[j][W-i-1])
			}
		}
		for i := 0; i < W; i++ {
			T[i] = U[i]
		}
		H, W = W, H
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
