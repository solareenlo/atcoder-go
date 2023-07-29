package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

var SS [500]string
var mind, dif [500]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscan(in, &Q, &SS[0])
	rebuild()
	ad := 0
	for Q > 0 {
		Q--
		var s string
		var y, z int
		fmt.Fscan(in, &s, &y, &z)
		y--
		if s[0] == '(' || s[0] == ')' {
			var i int
			for i = 0; i < 500; i++ {
				if y <= len(SS[i]) {
					break
				}
				y -= len(SS[i])
			}
			SS[i] = SS[i][:y] + s + SS[i][y:]
			ad++
			if ad >= 1000 {
				rebuild()
				ad = 0
			} else {
				rebuild2(i)
			}
		}
		if s[0] == 'D' {
			var i int
			for i = 0; i < 500; i++ {
				if y < len(SS[i]) {
					break
				}
				y -= len(SS[i])
			}
			SS[i] = SS[i][:y] + SS[i][y+1:]
			rebuild2(i)
		}
		if s[0] == 'Q' {
			var iy, iz int
			for iy = 0; iy < 500; iy++ {
				if y < len(SS[iy]) {
					break
				}
				y -= len(SS[iy])
			}
			for iz = 0; iz < 500; iz++ {
				if z <= len(SS[iz]) {
					break
				}
				z -= len(SS[iz])
			}

			if iy == iz {
				p := dodo(iy, y, z)
				fmt.Println(p.x + (p.y + p.x))
			} else {
				di := 0
				mi := 0
				p := dodo(iy, y, len(SS[iy]))
				mi = p.x
				di = p.y + p.x
				for i := iy + 1; i < iz; i++ {
					p = pair{mind[i], dif[i]}
					if p.x > di {
						mi += p.x - di
						di += p.y + (p.x - di)
					} else {
						di += p.y
					}
				}
				p = dodo(iz, 0, z)
				if p.x > di {
					mi += p.x - di
					di += p.y + (p.x - di)
				} else {
					di += p.y
				}
				fmt.Println(mi + di)
			}
		}
	}
}

func rebuild() {
	var s string
	for i := 0; i < 500; i++ {
		s += SS[i]
	}
	for i := 0; i < 500; i++ {
		if len(s) <= i*500 {
			SS[i] = ""
		} else if len(s) >= (i+1)*500 {
			SS[i] = s[i*500 : i*500+500]
		} else {
			SS[i] = s[i*500:]
		}
		rebuild2(i)
	}
}

func rebuild2(cur int) {
	p := dodo(cur, 0, len(SS[cur]))
	mind[cur] = p.x
	dif[cur] = p.y
}

func dodo(cur, y, z int) pair {
	mi := 0
	dif := 0
	for i := y; i < z; i++ {
		if SS[cur][i] == '(' {
			dif++
		} else {
			dif--
			mi = max(mi, -dif)
		}
	}
	return pair{mi, dif}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(a *[]int, index int, value int) []int {
	n := len(*a)
	if index < 0 {
		index = (index%n + n) % n
	}
	switch {
	case index == n:
		return append(*a, value)

	case index < n:
		*a = append((*a)[:index+1], (*a)[index:]...)
		(*a)[index] = value
		return *a

	case index < cap(*a):
		*a = (*a)[:index+1]
		for i := n; i < index; i++ {
			(*a)[i] = 0
		}
		(*a)[index] = value
		return *a

	default:
		b := make([]int, index+1)
		if n > 0 {
			copy(b, *a)
		}
		b[index] = value
		return b
	}
}
