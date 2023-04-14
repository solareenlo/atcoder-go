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

	var qn int
	fmt.Fscan(in, &qn)

	for qn > 0 {
		qn--
		var ei int
		var s string
		fmt.Fscan(in, &s, &ei)
		ci := int(s[0] - 'a')
		if ei == 0 {
			if ci == 0 {
				fmt.Fprintln(out, "a")
			} else {
				fmt.Fprintln(out, "aa")
			}
			continue
		}

		k := 0
		for (ei & 1) == 0 {
			k++
			ei >>= 1
		}
		if ei > 25 {
			fmt.Fprintln(out, -1)
			continue
		}

		for k > 0 && ei < ci && (ei<<1) < 25 {
			k--
			ei <<= 1
		}

		for i := 0; i < ci && ei > 0; i, ei = i+1, ei-1 {
			fmt.Fprint(out, string('a'+i))
		}
		for i := 0; i < k; i++ {
			fmt.Fprint(out, string('a'+ci))
		}
		for i := ci + 1; i < 26 && ei > 0; i, ei = i+1, ei-1 {
			fmt.Fprint(out, string('a'+i))
		}
		fmt.Fprintln(out)
	}
}
