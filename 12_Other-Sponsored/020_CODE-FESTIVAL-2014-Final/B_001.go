package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	ans := 0
	for i := 0; i < len(s); i++ {
		n := int(s[i] - '0')
		if i%2 != 0 {
			ans -= n
		} else {
			ans += n
		}
	}
	fmt.Println(ans)
}
