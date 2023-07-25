package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	str := strings.Split(getLine(), " ")
	S := make(map[string]struct{})
	for _, s := range str {
		x := false
		t := ""
		for _, a := range s {
			if a == '@' {
				x = true
				if t != "" {
					S[t] = struct{}{}
					t = ""
				}
			} else if x {
				t += string(a)
			}
		}
		if t != "" {
			S[t] = struct{}{}
		}
	}

	keys := make([]string, 0, len(S))
	for k := range S {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, i := range keys {
		fmt.Println(i)
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
