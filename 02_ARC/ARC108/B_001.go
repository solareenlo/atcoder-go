package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	t := make([]byte, 0)
	for i := range s {
		t = append(t, s[i])
		if len(t) >= 3 && t[len(t)-3] == 'f' && t[len(t)-2] == 'o' && t[len(t)-1] == 'x' {
			t = t[:len(t)-3]
		}
	}
	fmt.Println(len(t))
}
