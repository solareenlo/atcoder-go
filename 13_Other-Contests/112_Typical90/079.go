package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)

	var A [101][101]int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var B int
			fmt.Fscan(in, &B)
			A[i][j] -= B
		}
	}
	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			ans += abs(A[i][j])
			A[i][j+1] -= A[i][j]
			A[i+1][j] -= A[i][j]
			A[i+1][j+1] -= A[i][j]
		}
	}
	if A[H-1][W-1] != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
		fmt.Println(ans)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
