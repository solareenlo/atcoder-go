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

	var s string
	fmt.Fscan(in, &s)

	q := list.New()
	rev := false
	for i := range s {
		if s[i] == 'R' {
			rev = !rev
		} else if rev {
			q.PushFront(s[i])
		} else {
			q.PushBack(s[i])
		}
	}

	if rev {
		q2 := list.New()
		for q.Len() > 0 {
			c := q.Remove(q.Back()).(byte)
			q2.PushBack(c)
		}
		q = q2
	}

	t := list.New()
	for front := q.Front(); front != nil; front = front.Next() {
		if t.Len() != 0 && t.Back().Value == front.Value {
			t.Remove(t.Back())
		} else {
			t.PushBack(front.Value)
		}
	}
	for front := t.Front(); front != nil; front = front.Next() {
		fmt.Print(string(front.Value.(byte)))
	}
	fmt.Println()
}
