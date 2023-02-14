package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	ans := make(map[string]bool)
	for n > 0 {
		n--
		var S string
		fmt.Fscan(in, &S)
		x := strings.Index(S, "(")
		s := strings.Split(S, "")
		if x == -1 {
			for len(s) < 2000 {
				s = append(s, "0")
			}
		} else {
			t := make([]string, len(s[x+1:]))
			copy(t, s[x+1:])
			t = t[:len(t)-1]
			s = s[0:x]
			for len(s) < 2000 {
				s = append(s, t...)
			}
			s = s[:2000]
		}
		if s[1999] == "9" {
			for i := 1999; i >= 0; i-- {
				if s[i] == "9" {
					s[i] = "0"
				} else {
					if s[i] == "." {
						s[i-1] = string(s[i-1][0] + 1)
					} else {
						s[i] = string(s[i][0] + 1)
					}
					break
				}
			}
		}
		ans[strings.Join(s, "")] = true
	}
	fmt.Println(len(ans))
}
