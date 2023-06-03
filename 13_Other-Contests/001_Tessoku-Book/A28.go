package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	ans := 0
	for n > 0 {
		n--
		var t string
		var a int
		fmt.Fscan(in, &t, &a)
		if t == "+" {
			ans += a
		} else if t == "-" {
			ans -= a
			ans += 10000
		} else {
			ans *= a
		}
		ans %= 10000
		fmt.Fprintln(out, ans)
	}
}
