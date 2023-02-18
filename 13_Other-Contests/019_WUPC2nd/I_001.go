package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	const OFFSET = 1100
	var a1 [2200][2200]int
	var a2 [2200][2200]int
	var a3 [2200][2200]int
	for n > 0 {
		n--
		var TYPE, x, y, sz int
		fmt.Fscan(in, &TYPE, &x, &y, &sz)
		x += OFFSET
		y += OFFSET
		switch TYPE {
		case 1:
			a1[y+sz-1][x-sz+1]++
			a1[y-sz][x+sz]--
			for i := 1; i < sz; i++ {
				a1[y+sz-1-i][x-sz+1]++
				a1[y+sz-1][x-sz+1+i]++
				a1[y-sz+i][x+sz]--
				a1[y-sz][x+sz-i]--
			}
		case 2:
			for i := 0; i < sz; i++ {
				a2[y+i][x-sz+1]++
				a2[y+i][x+1]--
				a2[y-i-1][x+1]--
				a2[y-i-1][x+sz+1]++
			}
		default:
			for i := 0; i < sz; i++ {
				a2[y+i][x-sz+1]++
				a2[y+i][x+1] -= 2*(sz-i) - 1
				a2[y-i-1][x+1] -= 2*(sz-i) - 1
				a2[y-i-1][x+sz+1]++
			}
			for i := 0; i < sz-1; i++ {
				a3[y+i][x-sz+2] += 2
				a3[y+i][x+1-i] -= 2
				a3[y-i-1][x+2+i] += 2
				a3[y-i-1][x+sz+1] -= 2
			}
		}
	}

	for y := 0; y < 2200; y++ {
		for x := 0; x < 2200-1; x++ {
			a3[y][x+1] += a3[y][x]
		}
	}
	for y := 0; y < 2200; y++ {
		for x := 0; x < 2200-1; x++ {
			a2[y][x] += a3[y][x]
		}
	}

	for y := 0; y < 2200; y++ {
		for x := 0; x < 2200-1; x++ {
			a2[y][x+1] += a2[y][x]
		}
	}
	for y := 0; y < 2200; y++ {
		for x := 0; x < 2200-1; x++ {
			a1[y][x] += a2[y][x]
		}
	}

	for y := 0; y < 2200; y++ {
		for d := 0; d < 2200; d++ {
			if y-d-1 >= 0 {
				a1[y-d-1][d+1] += a1[y-d][d]
			}
		}
	}
	for x := 0; x < 2200; x++ {
		for d := 0; d < 2200; d++ {
			if x > 0 && x+d+1 < 2200 {
				a1[2200-2-d][x+d+1] += a1[2200-1-d][x+d]
			}
		}
	}

	ans := 0
	for y := 0; y < 2200; y++ {
		for x := 0; x < 2200; x++ {
			ans = max(ans, a1[y][x])
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
