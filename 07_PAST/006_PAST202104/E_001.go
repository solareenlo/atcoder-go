package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	var S string
	fmt.Fscan(in, &N, &S)

	q := list.New()
	s := make([]int, 0)
	for i := 1; i <= N; i++ {
		if S[i-1] == 'L' {
			q.PushFront(i)
		} else if S[i-1] == 'R' {
			q.PushBack(i)
		} else {
			a := int(S[i-1] - 'A')
			if a < 3 {
				for j := 0; j < a && q.Len() > 0; j++ {
					s = append(s, q.Front().Value.(int))
					q.Remove(q.Front())
				}
				if q.Len() == 0 {
					fmt.Fprintln(out, "ERROR")
				} else {
					fmt.Fprintln(out, q.Front().Value.(int))
					q.Remove(q.Front())
				}
				for len(s) != 0 {
					q.PushFront(s[len(s)-1])
					s = s[:len(s)-1]
				}
			} else {
				a -= 3
				for j := 0; j < a && q.Len() > 0; j++ {
					s = append(s, q.Back().Value.(int))
					q.Remove(q.Back())
				}
				if q.Len() == 0 {
					fmt.Fprintln(out, "ERROR")
				} else {
					fmt.Fprintln(out, q.Back().Value.(int))
					q.Remove(q.Back())
				}
				for len(s) != 0 {
					q.PushBack(s[len(s)-1])
					s = s[:len(s)-1]
				}
			}
		}
	}
}
