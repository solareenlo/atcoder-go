package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]string, 1000)
	for i := 1; i <= 1000; i++ {
		s[i-1] = strconv.Itoa(i)
	}
	sort.Strings(s)
	for i := 0; i < 1000; i++ {
		fmt.Println(s[i])
	}
}
