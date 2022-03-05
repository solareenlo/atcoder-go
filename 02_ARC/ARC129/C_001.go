package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	mo := 1
	jg := false
	for n > 0 {
		k := int(math.Floor(math.Sqrt(2.0*float64(n)+1.0/4.0) - 1/2.0))
		if jg {
			fmt.Fprint(out, mo)
			mo *= 10
			mo %= 7
		}
		for i := 0; i < k; i++ {
			fmt.Fprint(out, 7)
			mo *= 10
			mo %= 7
		}
		jg = true
		n -= (k*k + k) / 2
	}
}
