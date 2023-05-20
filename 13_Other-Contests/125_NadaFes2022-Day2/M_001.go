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

	var S string
	fmt.Fscan(in, &S)
	N := len(S)
	var L int
	fmt.Fscan(in, &L)

	cnt := (L + 1) / 2
	c := 0
	for i := 0; i < N; i++ {
		if S[i] == '(' {
			c++
		}
	}
	if c < cnt {
		for k := 0; k < N-L+1; k++ {
			fmt.Fprintln(out, -1)
		}
		return
	}

	posA := make([]int, 0)
	posB := make([]int, 0)
	for i := 0; i < N; i++ {
		if S[i] == '(' {
			c := len(posA)
			posA = append(posA, i-2*c)
			posB = append(posB, i-c)
		}
	}

	k := 0
	posL := 0

	rqA := make([]Func, N*2)
	rqB := make([]Func, N*2)
	var sum_rqA, sum_rqB Func

	var add_node func(int)
	add_node = func(i int) {
		rqA[N+posA[i]].plus(Func{1, posA[i]})
		if k-posL*2 <= posA[i] {
			sum_rqA.plus(Func{1, posA[i]})
		}
		rqB[N+posB[i]].plus(Func{1, posB[i]})
		if posB[i] < k-posL {
			sum_rqB.plus(Func{1, posB[i]})
		}
	}
	var rem_node func(int)
	rem_node = func(i int) {
		rqA[N+posA[i]].minus(Func{1, posA[i]})
		if k-posL*2 <= posA[i] {
			sum_rqA.minus(Func{1, posA[i]})
		}
		rqB[N+posB[i]].minus(Func{1, posB[i]})
		if posB[i] < k-posL {
			sum_rqB.minus(Func{1, posB[i]})
		}
	}
	var increment_k func()
	increment_k = func() {
		sum_rqA.minus(rqA[N+k-posL*2])
		sum_rqB.plus(rqB[N+k-posL])
		k++
	}
	var increment_posL func()
	increment_posL = func() {
		rem_node(posL)
		add_node(posL + cnt)
		sum_rqA.plus(rqA[N+k-posL*2-1])
		sum_rqA.plus(rqA[N+k-posL*2-2])
		sum_rqB.minus(rqB[N+k-posL-1])
		posL++
	}
	var decrement_posL func()
	decrement_posL = func() {
		posL--
		sum_rqA.minus(rqA[N+k-posL*2-1])
		sum_rqA.minus(rqA[N+k-posL*2-2])
		sum_rqB.plus(rqB[N+k-posL-1])
		add_node(posL)
		rem_node(posL + cnt)
	}

	for i := 0; i < cnt; i++ {
		add_node(i)
	}

	var calc_score func() int
	calc_score = func() int {
		cA := sum_rqA
		cB := sum_rqB
		ans := (cA.c - cA.x*(k-posL*2)) + (-cB.c + cB.x*(k-posL))
		return ans
	}

	for kk := 0; kk < N-L+1; kk++ {
		if kk != 0 {
			increment_k()
		}
		score := calc_score()
		for posL+cnt != len(posA) {
			increment_posL()
			new_score := calc_score()
			if score <= new_score {
				decrement_posL()
				break
			}
			score = new_score
		}
		fmt.Fprintln(out, score)
	}
}

type Func struct {
	x, c int
}

func (l *Func) plus(r Func) {
	l.x += r.x
	l.c += r.c
}

func (l *Func) minus(r Func) {
	l.x -= r.x
	l.c -= r.c
}
