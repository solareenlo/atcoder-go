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

	fr := list.New()
	ba := list.New()

	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var c string
		fmt.Fscan(in, &c)
		if c == "A" {
			var s string
			fmt.Fscan(in, &s)
			ba.PushBack(s)
		} else if c == "B" {
			var s string
			fmt.Fscan(in, &s)
			ba.PushFront(s)
		} else if c == "C" {
			fr.Remove(fr.Front())
		} else {
			fmt.Fprintln(out, fr.Front().Value)
		}
		for ba.Len() > fr.Len() {
			fr.PushBack(ba.Front().Value)
			ba.Remove(ba.Front())
		}
	}
}
