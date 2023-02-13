package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for {
		var tmp1 string
		fmt.Fscan(in, &tmp1)
		if tmp1 == "0" {
			break
		}
		tmp2 := strings.Split(tmp1, "/")
		A := tmp2[0]
		a, _ := strconv.Atoi(A)
		B := tmp2[1]
		b, _ := strconv.Atoi(B)
		C := tmp2[2]
		c, _ := strconv.Atoi(C)
		var tmp3 string
		fmt.Fscan(in, &tmp3)
		tmp4 := strings.Split(tmp3, ":")
		D := tmp4[0]
		d, _ := strconv.Atoi(D)
		E := tmp4[1]
		e, _ := strconv.Atoi(E)
		F := tmp4[2]
		f, _ := strconv.Atoi(F)
		var n string
		fmt.Fscan(in, &n)
		ti := 0
		for i := 0; i < len(n); i++ {
			ti *= 2
			if n[i] == '1' {
				ti++
			}
		}
		day := ti / 86400
		ti %= 86400
		f += ti % 60
		ti /= 60
		if f >= 60 {
			e++
			f -= 60
		}
		e += ti % 60
		ti /= 60
		if e >= 60 {
			d++
			e -= 60
		}
		d += ti % 24
		if d >= 24 {
			day++
			d -= 24
		}
		for day > 0 {
			day--
			c++
			if c > 31 || ((b == 4 || b == 6 || b == 9 || b == 11) && c > 30) || (b == 2 && c > 29) || (b == 2 && c > 28 && (a%4 != 0 || (a%400 != 0 && a%100 == 0))) {
				b++
				c = 1
			}
			if b > 12 {
				a++
				b = 1
			}
		}
		fmt.Fprintf(out, "%d/%02d/%02d %02d:%02d:%02d\n", a, b, c, d, e, f)
	}
}
