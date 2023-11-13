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

	var h, w int
	fmt.Fscan(in, &h, &w)
	s := make([]string, h)
	t := make([]string, h)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}
	for i := range t {
		fmt.Fscan(in, &t[i])
	}

	for i := 0; i < h; i++ {
		slice := strings.Split(s[i], "")
		sort.Strings(slice)
		s[i] = strings.Join(slice, "")
		slice = strings.Split(t[i], "")
		sort.Strings(slice)
		t[i] = strings.Join(slice, "")
		if s[i] != t[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
