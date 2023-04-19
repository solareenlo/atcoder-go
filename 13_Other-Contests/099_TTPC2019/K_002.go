package main

import (
	"bufio"
	"fmt"
	"os"
)

var bp [1 << 16]int

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAXN = 100000

	var n int
	var s, t string
	fmt.Fscan(in, &n, &s, &t)

	BPinit()

	i, j := 0, 0
	m := uint(0x8000000000000000)
	var bs, bt [MAXN]uint
	for i < n {
		if s[i] == '1' {
			bs[j] |= m
		}
		if t[i] == '1' {
			bt[j] |= m
		}
		m >>= 1
		if m == 0 {
			j++
			m = 0x8000000000000000
		}
		i++
	}

	ans := 0
	for i := 0; ; i++ {
		c1 := 0
		j := 0
		r := i
		for r > 0 {
			if r < 64 {
				c1 += bit_count((bs[j] ^ bt[j]) >> (64 - r))
				break
			} else {
				c1 += bit_count(bs[j] ^ bt[j])
				r -= 64
				j++
			}
		}

		c2 := 0
		c2 += bit_count((bs[j] ^ bt[j]) << r)
		r = n - i - (64 - r)
		for r > 0 {
			if r < 64 {
				j++
				c2 += bit_count((bs[j] ^ bt[j]) >> (64 - r))
				break
			} else {
				j++
				c2 += bit_count(bs[j] ^ bt[j])
				r -= 64
			}
		}

		if c1 >= c2 {
			ans = i
			break
		}
		j = (n - 1) / 64
		r = (n - 1) % 64
		l := (bs[j] << (r)) & 0x8000000000000000
		bs[j] &= ^(0x8000000000000000 >> r)
		bs[j] >>= 1
		for ; j > 0; j-- {
			bs[j] |= (bs[j-1] & 1) << 63
			bs[j-1] >>= 1
		}
		bs[0] |= l
	}
	fmt.Println(ans)
}

func BPinit() {
	for i := 0; i < 1<<16; i++ {
		t := i
		for t > 0 {
			if (t & 1) != 0 {
				bp[i]++
			}
			t >>= 1
		}
	}
}

func bit_count(t uint) int {
	c := 0
	for i := 0; i < 4; i++ {
		c += bp[(t>>(i*16))&0xFFFF]
	}
	return c
}
