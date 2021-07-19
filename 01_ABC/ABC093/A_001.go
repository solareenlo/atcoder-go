package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := make([]string, 1)
	fmt.Scan(&s[0])
	s = strings.Split(s[0], "")
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	if strings.Join(s, "") == "abc" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
