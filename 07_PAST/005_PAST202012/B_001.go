package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var S string
	fmt.Fscan(in, &n, &S)
	s := strings.Split(S, "")

	t := make([]string, 0)
	for _, x := range s {
		for j := 0; j < len(t); j++ {
			if t[j] == x {
				t[j] = "?"
			}
		}
		t = append(t, x)
	}

	for _, x := range t {
		if x != "?" {
			fmt.Fprint(out, x)
		}
	}
}
