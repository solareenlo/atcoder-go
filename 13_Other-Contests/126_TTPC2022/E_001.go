package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A := getLine()
	B := getLine()
	Q := getInt()
	var f func(byte) int
	f = func(c byte) int {
		if 'a' <= c && c <= 'z' {
			return int(c - 'a')
		}
		if 'A' <= c && c <= 'Z' {
			return int(c-'A') + 26
		}
		return 52
	}
	B = reverseString(B)
	tb1 := make([][]int, len(A)+1)
	for i := range tb1 {
		tb1[i] = make([]int, 53)
	}
	tb2 := make([][]int, len(B)+1)
	for i := range tb2 {
		tb2[i] = make([]int, 53)
	}
	tb := make([]int, 53)
	for i := range tb {
		tb[i] = -1
	}
	for i := len(A) - 1; i >= 0; i-- {
		copy(tb1[i+1], tb)
		tb[f(A[i])] = i + 1
	}
	copy(tb1[0], tb)
	for i := range tb {
		tb[i] = -1
	}
	for i := len(B) - 1; i >= 0; i-- {
		copy(tb2[i+1], tb)
		tb[f(B[i])] = i + 1
	}
	copy(tb2[0], tb)
	for Q > 0 {
		Q--
		s := getLine()
		pos1, pos2 := 0, 0
		for j := 0; pos1 < len(s); pos1++ {
			j = tb1[j][f(s[pos1])]
			if j == -1 {
				break
			}
		}
		s = reverseString(s)
		for j := 0; pos2 < len(s); pos2++ {
			j = tb2[j][f(s[pos2])]
			if j == -1 {
				break
			}
		}
		if pos1+pos2 < len(s) {
			fmt.Println(-1)
			continue
		}
		ans := len(s)
		for i := pos1; i >= 0; i-- {
			c := len(s) - i
			if c <= pos2 {
				ans = min(ans, abs(c-i))
			}
		}
		if ans == len(s) {
			fmt.Println(-1)
			continue
		}
		fmt.Println(ans)
	}
}

var reader = bufio.NewReader(os.Stdin)

func getLine() string {
	line, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return line[:len(line)-1]
}

func getInt() int {
	tmp := getLine()
	i, e := strconv.Atoi(tmp)
	if e != nil {
		panic(e)
	}
	return i
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
