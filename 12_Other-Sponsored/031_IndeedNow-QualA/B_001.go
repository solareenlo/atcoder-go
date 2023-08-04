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

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		slice := strings.Split(s, "")
		sort.Strings(slice)
		s = strings.Join(slice, "")
		if s == "ddeeinnow" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
