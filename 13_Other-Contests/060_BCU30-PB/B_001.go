package main

import (
	"bufio"
	"fmt"
	"os"
)

var dx = []int{1, 2, 2, 1, -1, -2, -2, -1}
var dy = []int{-2, -1, 1, 2, 2, 1, -1, -2}

func main() {
	var s [9]string
	for i := 0; i < 9; i++ {
		s[i] = readline()
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := j + 1; k < 9; k++ {
				if s[i][j] == s[i][k] || s[j][i] == s[k][i] {
					fmt.Println("No")
					return
				}
			}
			for k := 0; k < 8; k++ {
				nx := i + dx[k]
				ny := j + dy[k]
				if ny < 0 || ny >= 9 || nx < 0 || nx >= 9 {
					continue
				}
				if s[i][j] == s[nx][ny] {
					fmt.Println("No")
					return
				}
			}
		}
	}

	fmt.Println("Yes")
}

var reader = bufio.NewReaderSize(os.Stdin, 1000000)
var writer = bufio.NewWriter(os.Stdout)

func readline() string {
	buf := make([]byte, 0)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		buf = append(buf, line...)
		if !isPrefix {
			break
		}
	}
	return string(buf)
}
