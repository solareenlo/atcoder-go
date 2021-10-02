package main

import (
	"fmt"
	"sort"
)

func main() {
	var ss, tt string
	fmt.Scan(&ss, &tt)

	s, t := []byte(ss), []byte(tt)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})

	if string(s) < string(t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
