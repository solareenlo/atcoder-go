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

	cnt := [10]int{}
	for i := 0; i < n; i++ {
		var p, v string
		fmt.Fscan(in, &p, &v)
		if v[0] == 'A' && cnt[p[0]-'A'] == 0 {
			cnt[p[0]-'A'] = i + 1
		}
	}

	for i := 0; i < 6; i++ {
		fmt.Println(cnt[i])
	}
}
