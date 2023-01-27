package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	slice := strings.Split(s, "")
	sort.Strings(slice)
	s = strings.Join(slice, "")

	if s[0] != s[1] {
		fmt.Printf("%c\n", s[0])
	} else if s[1] != s[2] {
		fmt.Printf("%c\n", s[2])
	} else {
		fmt.Println(-1)
	}
}
