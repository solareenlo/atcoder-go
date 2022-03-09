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
	var s string
	fmt.Fscan(in, &n, &s)

	s = strings.ReplaceAll(s, "A", "BB")
	s = strings.ReplaceAll(s, "BB", "A")
	fmt.Fprintln(out, s)
}
