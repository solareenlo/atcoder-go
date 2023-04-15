package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	dark := 0
	deep := 1
	for n > 0 {
		n--
		var c string
		var a int
		fmt.Fscan(in, &c, &a)
		if c == "+" {
			dark += a
		} else if c == "*" && a >= 1 {
			deep *= a
		}
	}
	fmt.Println(dark * deep)
}
