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

	var h, w int
	fmt.Fscan(in, &h, &w)
	for i := 0; i < h; i++ {
		tmp := strings.Repeat(".", w)
		s := strings.Split(tmp, "")
		for j := 0; j < w; j++ {
			var a int
			fmt.Fscan(in, &a)
			if a != 0 {
				s[j] = string(a + 64)
			}
		}
		fmt.Fprintln(out, strings.Join(s, ""))
	}
}
