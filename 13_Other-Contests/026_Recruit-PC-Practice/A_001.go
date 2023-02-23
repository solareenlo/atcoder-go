package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n int
		fmt.Fscan(in, &n)
		p := make([]string, 0)
		for n > 0 {
			n--
			var tmp string
			fmt.Fscan(in, &tmp)
			p = append(p, tmp)
		}
		cnt := 0
		for len(p) != 1 && cnt < int(1e6) {
			tmp := p[0]
			p = p[1:]
			c := p[0][0]
			p[0] = p[0][1:]
			if len(p[0]) == 0 {
				p = p[1:]
			}
			if idx := strings.IndexByte(tmp, c); idx != -1 {
				tmp = tmp[:idx] + tmp[idx+1:]
			} else {
				tmp += string(c)
			}
			if len(tmp) != 0 {
				p = append(p, tmp)
			}
			cnt++
		}
		if cnt >= int(1e6) {
			fmt.Println(-1)
		} else {
			fmt.Println(cnt)
		}
	}
}
