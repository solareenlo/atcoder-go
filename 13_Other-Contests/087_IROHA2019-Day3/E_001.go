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
	ans := 0
	now := 0
	f := true
	for i := 0; i < n; i++ {
		var c string
		fmt.Fscan(in, &c)
		if c == "/" {
			if !f && now == 0 {
				ans++
			}
			if !f {
				now = 0
			}
			f = true
			now++
		} else {
			f = false
			now--
		}
	}
	if !f && now == 0 {
		ans++
	}
	fmt.Println(ans)
}
