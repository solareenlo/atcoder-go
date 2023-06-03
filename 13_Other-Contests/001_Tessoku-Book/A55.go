package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAX = 114514

var zac int
var zat []int
var bit_table [MAX]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var query [MAX][2]int

	var Q int
	fmt.Fscan(in, &Q)
	zat = make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &query[i][0], &query[i][1])
		zat[i] = query[i][1]
	}
	sort.Ints(zat)
	zac = 1
	for i := 1; i < Q; i++ {
		if zat[zac-1] != zat[i] {
			zat[zac] = zat[i]
			zac++
		}
	}
	for i := 0; i < Q; i++ {
		id := zaq(query[i][1])
		switch query[i][0] {
		case 1:
			bit_add(id, 1)
		case 2:
			bit_add(id, -1)
		case 3:
			{
				less := bit_sum(id - 1)
				if bit_sum(zac-1)-less <= 0 {
					fmt.Println(-1)
				} else {
					if bit_sum(id)-less > 0 {
						fmt.Println(zat[id])
					} else {
						no := id
						yes := zac - 1
						for no+1 < yes {
							m := no + (yes-no)/2
							if bit_sum(m)-less > 0 {
								yes = m
							} else {
								no = m
							}
						}
						fmt.Println(zat[yes])
					}
				}
			}
		}
	}
}

func zaq(q int) int {
	l := 0
	r := zac - 1
	for l <= r {
		m := l + (r-l)/2
		if zat[m] == q {
			return m
		} else if zat[m] < q {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func bit_add(pos, value int) {
	pos++
	for pos <= MAX {
		bit_table[pos-1] += value
		pos += pos & (-pos)
	}
}

func bit_sum(pos int) int {
	sum := 0
	pos++
	for pos > 0 {
		sum += bit_table[pos-1]
		pos -= pos & (-pos)
	}
	return sum
}
