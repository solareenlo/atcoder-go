package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 300300

var cnt, val [N * 30]int
var ch [N * 30][2]int
var tot int = 1

func pushup(x int) {
	cnt[x] = cnt[ch[x][0]] + cnt[ch[x][1]]
	if cnt[x] == 0 {
		val[x] = 0
	} else if cnt[x] == 1 {
		val[x] = val[ch[x][0]] + val[ch[x][1]]
	} else if cnt[ch[x][0]] == 1 && cnt[ch[x][1]] == 1 {
		val[x] = val[ch[x][0]] ^ val[ch[x][1]]
	} else {
		val[x] = 2e9
		if cnt[ch[x][0]] >= 2 {
			val[x] = min(val[x], val[ch[x][0]])
		}
		if cnt[ch[x][1]] >= 2 {
			val[x] = min(val[x], val[ch[x][1]])
		}
	}
}

func ins(dep, x, v int) {
	if dep < 0 {
		cnt[x]++
		if cnt[x] == 1 {
			val[x] = v
		} else {
			val[x] = 0
		}

		return
	}
	v0 := (v >> dep) & 1
	if ch[x][v0] == 0 {
		tot++
		ch[x][v0] = tot
	}
	ins(dep-1, ch[x][v0], v)
	pushup(x)
}

func del(dep, x, v int) {
	if dep < 0 {
		cnt[x]--
		if cnt[x] == 1 {
			val[x] = v
		} else {
			val[x] = 0
		}
		return
	}
	v0 := (v >> dep) & 1
	del(dep-1, ch[x][v0], v)
	pushup(x)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var op int
		fmt.Fscan(in, &op)
		if op == 1 {
			var x int
			fmt.Fscan(in, &x)
			ins(29, 1, x)
		} else if op == 2 {
			var x int
			fmt.Fscan(in, &x)
			del(29, 1, x)
		} else {
			fmt.Fprintln(out, val[1])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
