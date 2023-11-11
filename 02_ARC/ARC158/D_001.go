package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var solve func()
	solve = func() {
		var n, Mod int
		fmt.Fscan(in, &n, &Mod)

		var Pow func(int, int) int
		Pow = func(a, n int) int {
			res := 1
			for n > 0 {
				if n%2 == 1 {
					res = res * a % Mod
				}
				a = a * a % Mod
				n /= 2
			}
			return res
		}

		var Rand func(int) int
		Rand = func(n int) int {
			return r.Intn(n)
		}

		for {
			a := 1
			b := Mod - 1
			c := Rand(Mod - 2)
			for c == a || c == b {
				c = Rand(Mod - 2)
			}

			a1 := a
			var b1 int
			if (n & 1) != 0 {
				b1 = b
			} else {
				b1 = a
			}
			c1 := Pow(c, n)
			a2 := a1 * a1 % Mod
			b2 := b1 * b1 % Mod
			c2 := c1 * c1 % Mod
			a3 := a1 * a2 % Mod
			b3 := b1 * b2 % Mod
			c3 := c1 * c2 % Mod
			x := (a + b + c) * (a1 + b1 + c1) % Mod * (a2 + b2 + c2) % Mod
			y := (a3 + b3 + c3) % Mod

			if x == 0 || y == 0 {
				continue
			}

			t := y * Pow(x, Mod-2) % Mod

			a = a * t % Mod
			b = b * t % Mod
			c = c * t % Mod

			if a > b {
				a, b = b, a
			}
			if b > c {
				b, c = c, b
			}
			if a > b {
				a, b = b, a
			}
			if b > c {
				b, c = c, b
			}

			fmt.Fprintln(out, a, b, c)

			break
		}
	}

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		solve()
	}
}
