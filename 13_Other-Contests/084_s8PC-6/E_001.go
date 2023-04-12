package main

import (
	"bufio"
	"fmt"
	"os"
)

var X [209][]int
var col [209]int
var cnts int

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)
	var c [109]string
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &c[i])
		for j := 0; j < W; j++ {
			if c[i][j] == 'o' {
				X[i] = append(X[i], H+j)
				X[H+j] = append(X[H+j], i)
			}
		}
	}
	for i := 0; i < H+W; i++ {
		if len(X[i]) == 0 || col[i] >= 1 {
			continue
		}
		cnts++
		dfs(i)
	}
	flag := true
	if cnts >= 2 {
		flag = false
	}
	cnt1 := 0
	for i := 0; i < H+W; i++ {
		if len(X[i])%2 == 1 {
			cnt1++
		}
	}
	if cnt1 >= 4 {
		flag = false
	}

	if flag == true {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}
}

func dfs(pos int) {
	if col[pos] >= 1 {
		return
	}
	col[pos] = cnts
	for i := 0; i < len(X[pos]); i++ {
		dfs(X[pos][i])
	}
}
