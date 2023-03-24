package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, ans, a, Max int64
	fmt.Fscan(in, &n)
	ans += n
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &a)
		ans += a
		if a > Max {
			Max = a
		}
	}
	fmt.Println(ans + Max)
}
