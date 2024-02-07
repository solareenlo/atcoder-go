package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	s = sortString(s)

	ans := 0
	for i := 0; i <= 3162277; i++ {
		j := i * i
		t := strconv.Itoa(j)
		for len(t) < len(s) {
			t += "0"
		}
		t = sortString(t)
		if s == t {
			ans++
		}
	}
	fmt.Println(ans)
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
