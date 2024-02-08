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

	const N = 200005

	var n, m int
	fmt.Fscan(in, &n, &m)

	ans := n + 1
	var a, cnt [N]int
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
		cnt[a[i]]++
		if cnt[a[i]] > cnt[ans] {
			ans = a[i]
		} else if cnt[a[i]] == cnt[ans] {
			if a[i] < ans {
				ans = a[i]
			}
		}
		fmt.Fprintln(out, ans)
	}
}
