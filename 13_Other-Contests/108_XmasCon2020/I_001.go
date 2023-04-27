package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, l, n int
	var S string
	fmt.Fscan(in, &m, &l, &n, &S)

	equations := make([][]int, 0)
	s := n + 1

	var get_new func([]int) int
	get_new = func(v []int) int {
		k := s
		s++
		v = append([]int{k}, v...)
		equations = append(equations, v)
		return k
	}

	saved_not := make(map[int]int)

	var NOT func(int) int
	NOT = func(a int) int {
		if _, ok := saved_not[a]; !ok {
			tmp := make([]int, m)
			for i := range tmp {
				tmp[i] = a
			}
			saved_not[a] = get_new(tmp)
		}
		return saved_not[a]
	}

	var ret func(int)
	ret = func(a int) {
		b := NOT(a)
		v := make([]int, m)
		for i := range v {
			v[i] = b
		}
		v = append([]int{n}, v...)
		equations = append(equations, v)
	}

	var zero, one int

	var OR func(int, int) int
	OR = func(a, b int) int {
		if l < m {
			z := make([]int, 0)
			for j := 0; j < l-1; j++ {
				z = append(z, one)
			}
			for j := 0; j < m-l-1; j++ {
				z = append(z, zero)
			}
			z = append(z, a)
			z = append(z, b)
			return NOT(get_new(z))
		} else {
			z := make([]int, 0)
			for j := 0; j < m-2; j++ {
				z = append(z, one)
			}
			z = append(z, NOT(a))
			z = append(z, NOT(b))
			return get_new(z)
		}
	}

	var AND func(int, int) int
	AND = func(a, b int) int {
		a = NOT(a)
		b = NOT(b)
		c := OR(a, b)
		return NOT(c)
	}

	if m == (2*l - 1) {
		zero = 0
	} else {
		g := NOT(0)
		k := m / 2
		z := make([]int, 0)
		for i := 0; i < k; i++ {
			z = append(z, 0)
		}
		for i := 0; i < m-k; i++ {
			z = append(z, g)
		}
		if k >= l {
			zero = get_new(z)
		} else {
			zero = NOT(get_new(z))
		}
	}
	one = NOT(zero)

	if m == 1 {
		found := false
		for j := 0; j < n; j++ {
			var r0, r1 string
			for i := 0; i < (1 << n); i++ {
				if ((i >> j) & 1) != 0 {
					r1 += string(S[i])
				} else {
					r0 += string(S[i])
				}
			}
			if r0 != r1 && r0 == strings.Repeat(string(r0[0]), len(r0)) && r1 == strings.Repeat(string(r1[0]), len(r1)) {
				if r0[0] == '0' {
					ret(j)
				} else {
					ret(NOT(j))
				}
				found = true
				break
			}
		}
		if !found {
			fmt.Println(-1)
			return
		}
	} else {
		for i := 0; i < (1 << n); i++ {
			if m == (2*l-1) && S[i] == S[(1<<n)-1-i] {
				fmt.Println(-1)
				return
			}
		}
		val := zero
		wut := make([]int, 1)
		wut[0] = one
		for j := 0; j < n; j++ {
			nwut := make([]int, 1<<(j+1))
			for r := 0; r < (1 << (j + 1)); r++ {
				if ((r >> j) & 1) != 0 {
					nwut[r] = AND(wut[r^(1<<j)], j)
				} else {
					nwut[r] = AND(wut[r], NOT(j))
				}
			}
			wut = nwut
		}
		for i := 0; i < (1 << n); i++ {
			if S[i] == '1' {
				val = OR(val, wut[i])
			}
		}
		ret(val)
	}

	fmt.Println(len(equations))
	for _, v := range equations {
		for _, a := range v {
			fmt.Printf("%d ", a)
		}
		fmt.Println()
	}
}
