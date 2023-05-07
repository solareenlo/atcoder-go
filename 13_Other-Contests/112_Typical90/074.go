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
	var s string
	fmt.Fscan(in, &s)

	ans := 0
	for i := 0; i < n; i++ {
		ans += int(s[i]-'a') * (1 << i)
	}
	fmt.Println(ans)
}
