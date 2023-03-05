package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const MOD = 1000000007
const EPS = 1e-12

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	ch := NewConvex_hull(false)
	x := make([]float64, N)
	y := make([]float64, N)
	r := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i], &y[i], &r[i])
	}
	d := make([]float64, M)
	u := make([]float64, M)
	rr := make([]float64, M)
	l := make([]float64, M)
	flag := make([]bool, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &l[i], &d[i], &rr[i], &u[i])
		rr[i] += l[i]
		u[i] += d[i]
		for j := 0; j < N; j++ {
			if l[i] <= x[j] && x[j] <= rr[i] && d[i] <= y[j] && y[j] <= u[i] {
				flag[i] = true
			}
			if (x[j]-l[i])*(x[j]-l[i])+(y[j]-u[i])*(y[j]-u[i]) <= r[j]*r[j] {
				flag[i] = true
			}
			if (x[j]-l[i])*(x[j]-l[i])+(y[j]-d[i])*(y[j]-d[i]) <= r[j]*r[j] {
				flag[i] = true
			}
			if (x[j]-rr[i])*(x[j]-rr[i])+(y[j]-d[i])*(y[j]-d[i]) <= r[j]*r[j] {
				flag[i] = true
			}
			if (x[j]-rr[i])*(x[j]-rr[i])+(y[j]-u[i])*(y[j]-u[i]) <= r[j]*r[j] {
				flag[i] = true
			}
			if l[i] <= x[j] && x[j] <= rr[i] {
				if math.Abs(u[i]-y[j]) <= r[j] {
					flag[i] = true
				}
				if math.Abs(d[i]-y[j]) <= r[j] {
					flag[i] = true
				}
			}
			if d[i] <= y[j] && y[j] <= u[i] {
				if math.Abs(l[i]-x[j]) <= r[j] {
					flag[i] = true
				}
				if math.Abs(rr[i]-x[j]) <= r[j] {
					flag[i] = true
				}
			}
		}
		if flag[i] {
			ch.Add_node(Coordinate{l[i], u[i]})
			ch.Add_node(Coordinate{l[i], d[i]})
			ch.Add_node(Coordinate{rr[i], u[i]})
			ch.Add_node(Coordinate{rr[i], d[i]})
		}
	}
	if len(ch.node) == 0 {
		fmt.Println(0)
		return
	}
	node := ch.solve()
	ans := 0.0
	K := len(node)
	for i := 0; i < K; i++ {
		ans += math.Sqrt((node[i].x-node[(i+K-1)%K].x)*(node[i].x-node[(i+K-1)%K].x) + (node[i].y-node[(i+K-1)%K].y)*(node[i].y-node[(i+K-1)%K].y))
	}
	fmt.Printf("%.20f\n", ans)
}

type Convex_hull struct {
	node []Coordinate
	ret  []Coordinate
	line bool
}

func NewConvex_hull(l bool) *Convex_hull {
	c := new(Convex_hull)
	c.node = make([]Coordinate, 0)
	c.ret = make([]Coordinate, 0)
	c.line = l
	return c
}

func (c *Convex_hull) Add_node(n Coordinate) {
	c.node = append(c.node, n)
}

func (c *Convex_hull) solve() []Coordinate {
	sort.Slice(c.node, func(i, j int) bool {
		return c.node[i].lessThan(c.node[j])
	})
	index := 0
	num := len(c.node)
	c.ret = make([]Coordinate, num*2)
	for i := 0; i < num; i++ {
		if c.line {
			for index > 1 && (c.ret[index-1].minus(c.ret[index-2])).det(c.node[i].minus(c.ret[index-1])) < -EPS {
				index--
			}
		} else {
			for index > 1 && (c.ret[index-1].minus(c.ret[index-2])).det(c.node[i].minus(c.ret[index-1])) < EPS {
				index--
			}
		}
		c.ret[index] = c.node[i]
		index++
	}
	box := index
	for i := num - 2; i >= 0; i-- {
		if c.line {
			for index > box && (c.ret[index-1].minus(c.ret[index-2])).det(c.node[i].minus(c.ret[index-1])) < -EPS {
				index--
			}
		} else {
			for index > box && (c.ret[index-1].minus(c.ret[index-2])).det(c.node[i].minus(c.ret[index-1])) < EPS {
				index--
			}
		}
		c.ret[index] = c.node[i]
		index++
	}
	resize(&(c.ret), index-1)
	return c.ret
}

func resize(a *[]Coordinate, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, Coordinate{0.0, 0.0})
		}
	}
}

type Coordinate struct {
	x, y float64
}

func (a Coordinate) lessThan(c Coordinate) bool {
	if a.y < c.y {
		return true
	}
	if a.y > c.y {
		return false
	}
	if a.x < c.x {
		return true
	}
	return false
}

func (a Coordinate) plus(c Coordinate) Coordinate {
	var box Coordinate
	box.x = a.x + c.x
	box.y = a.y + c.y
	return box
}

func (a Coordinate) minus(c Coordinate) Coordinate {
	var box Coordinate
	box.x = a.x - c.x
	box.y = a.y - c.y
	return box
}

func (c Coordinate) det(a Coordinate) float64 {
	return c.x*a.y - c.y*a.x
}
