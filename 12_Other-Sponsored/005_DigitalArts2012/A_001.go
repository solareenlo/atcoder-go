package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	S := getLine()
	S += " "
	s := strings.Split(S, "")
	n := getInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = getLine()
	}
	u := 0
	for i := 0; i < len(s); i++ {
		if s[i] == " " {
			t := s[u:i]
			for j := 0; j < n; j++ {
				if len(a[j]) == len(t) {
					f := true
					for k := 0; k < len(t); k++ {
						if a[j][k] != '*' && a[j][k] != t[k][0] {
							f = false
						}
					}
					if f {
						for k := 0; k < len(t); k++ {
							s[u+k] = "*"
						}
					}
				}
			}
			u = i + 1
		}
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Print(s[i])
	}
	fmt.Println()
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
