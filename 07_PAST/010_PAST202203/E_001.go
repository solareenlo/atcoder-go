package main

import "fmt"

func main() {
	var c [10]int

	var y, m, d int
	fmt.Scanf("%d/%d/%d", &y, &m, &d)
	for {
		for i := 0; i < 10; i++ {
			c[i] = 0
		}
		y1 := y
		x := y1 % 10
		c[x]++
		y1 = y1 / 10
		x = y1 % 10
		c[x]++
		y1 = y1 / 10
		x = y1 % 10
		c[x]++
		y1 = y1 / 10
		c[y1]++
		d1 := d
		x = d1 % 10
		c[x]++
		d1 = d1 / 10
		c[d1]++
		m1 := m
		x = m1 % 10
		c[x]++
		m1 = m1 / 10
		c[m1]++
		cnt := 0
		for i := 0; i < 10; i++ {
			if c[i] != 0 {
				cnt++
			}
		}
		if cnt <= 2 {
			break
		}
		d++
		if d > 22 {
			m++
			d = 1
		}
		if m > 12 {
			y++
			m = 1
		}
	}
	fmt.Printf("%d/", y)
	if m >= 10 {
		fmt.Printf("%d/", m)
	} else {
		fmt.Printf("0%d/", m)
	}
	if d >= 10 {
		fmt.Printf("%d\n", d)
	} else {
		fmt.Printf("0%d\n", d)
	}
}
