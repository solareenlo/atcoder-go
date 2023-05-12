package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	v := make([][]string, 500)
	var H, W, T int
	fmt.Fscan(in, &H, &W, &T)
	for i := 0; i < 500; i++ {
		var tmp string
		fmt.Fscan(in, &tmp)
		v[i] = strings.Split(tmp, "")
	}
	v = transpose(v)
	op(0, 70, v)
	op(70, 140, v)
	op(140, 210, v)
	op(210, 280, v)
	op(280, 390, v)
	op(390, 500, v)
	fmt.Fprintln(out, -1, -1)
	out.Flush()
}

func op(a, b int, v [][]string) {
	for i := 0; i < 500; i++ {
		for j := a; j < b; j++ {
			if (v[i][j] == "#" && j != 499 && v[i][j+1] == ".") || (j != 0 && v[i][j-1] == "." && v[i][j] == "#") || (v[i][j] == "#" && (j == b-1 || j == a)) {
				fmt.Fprintln(out, i, j)
				out.Flush()
				var T int
				fmt.Fscan(in, &T)
			}
		}
	}
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
