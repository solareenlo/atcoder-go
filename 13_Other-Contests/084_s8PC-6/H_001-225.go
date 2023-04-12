package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var cnt int
var q *list.List

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n int
		fmt.Fscan(in, &n)
		cnt = 0
		q = list.New()
		dfs(n)
		fmt.Fprintf(out, "%d ", cnt)
		for x := q.Front(); x != nil; x = x.Next() {
			fmt.Fprintf(out, "%d ", x.Value)
		}
		fmt.Fprintln(out)
	}
}

func dfs(n int) {
	if n == 1 {
		return
	} else if (n & 1) != 0 {
		dfs(n - 1)
		cnt++
		q.PushFront(cnt)
	} else {
		dfs(n >> 1)
		cnt++
		q.PushBack(cnt)
	}
}
