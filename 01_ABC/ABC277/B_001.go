package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	m := make(map[string]bool)
	r := regexp.MustCompile(`^([HDCS][A23456789TJQK])`)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		if r.MatchString(s) {
			m[s] = true
		}
	}

	if len(m) == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
