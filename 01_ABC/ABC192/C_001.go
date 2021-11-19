package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n string
	var k int
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		slice := strings.Split(n, "")
		sort.Strings(slice)
		s, _ := strconv.Atoi(strings.Join(slice, ""))
		sort.Sort(sort.Reverse(sort.StringSlice(slice)))
		t, _ := strconv.Atoi(strings.Join(slice, ""))
		n = strconv.Itoa(t - s)
	}
	fmt.Println(n)
}
