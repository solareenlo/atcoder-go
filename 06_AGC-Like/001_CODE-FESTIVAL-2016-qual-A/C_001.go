package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var S string
	var k int
	fmt.Fscan(in, &S, &k)
	s := strings.Split(S, "")

	for i := 0; i < len(s); i++ {
		tmp := (int('z'-s[i][0]) + 1) % 26
		if k >= tmp {
			k -= tmp
			s[i] = "a"
		}
	}

	k %= 26
	s[len(s)-1] = string(s[len(s)-1][0] + byte(k))
	fmt.Println(strings.Join(s, ""))
}
