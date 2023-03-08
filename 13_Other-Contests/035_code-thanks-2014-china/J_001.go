package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD = 1000000007
const INF = int(1e18)

var N, Q int
var A, pre []int
var fst [][]int
var stor []int
var m int
var T *Trie

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &Q)
	A = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	fst = make([][]int, N)
	for i := range fst {
		fst[i] = make([]int, 31)
	}
	for i := 0; i < 31; i++ {
		fst[N-1][i] = MOD
	}
	for i := N - 1; i >= 0; i-- {
		if i == N-1 {
			for j := 0; j < 3; j++ {
				fst[N-1][j] = MOD
			}
		} else {
			copy(fst[i], fst[i+1])
		}
		for j := 0; j < 31; j++ {
			if (A[i] & (1 << j)) == 0 {
				fst[i][j] = i
			}
		}
	}

	pre = make([]int, 0)
	pre = append(pre, 0)
	for _, t := range A {
		pre = append(pre, pre[len(pre)-1]^t)
	}

	stor = make([]int, 0)
	T = NewTrie()
	for i := 0; i < N+1; i++ {
		ind := 0
		if len(stor) != 0 {
			ind = stor[len(stor)-1]
		}
		tmp := T.insert(pre[i], 30, ind, i)
		stor = append(stor, tmp)
	}

	for i := 0; i < Q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		if i != 0 {
			l = (l + abs(m)) % N
			r = (r + abs(m)) % N
		} else {
			l--
			r--
		}
		m = query(l, r)
		fmt.Println(m)
	}
}

const MX = 4000000
const MXBIT = 31

type Trie struct {
	num int
	mx  [MX]int
	nex [MX][2]int
}

func NewTrie() *Trie {
	t := new(Trie)
	t.nex[0][0] = -1
	t.nex[0][1] = -1
	return t
}

func (T *Trie) insert(x, i, pre, label int) int {
	T.num++
	cur := T.num
	T.mx[cur] = label
	T.nex[cur][0] = -1
	T.nex[cur][1] = -1
	if pre != -1 {
		T.nex[cur][0] = T.nex[pre][0]
		T.nex[cur][1] = T.nex[pre][1]
	}
	if i == -1 {
		return cur
	}
	t := (x >> i) & 1
	tmp := -1
	if pre != -1 {
		tmp = T.nex[pre][t]
	}
	T.nex[cur][t] = T.insert(x, i-1, tmp, label)
	return cur
}

func (T *Trie) test(x, l, r int) int {
	cur := stor[r]
	for i := MXBIT - 1; i >= 0; i-- {
		t := (x >> i) & 1
		if T.nex[cur][t] == -1 || T.mx[T.nex[cur][t]] < l {
			t ^= 1
		}
		if T.nex[cur][t] == -1 {
			fmt.Print("WUT", i, cur, T.mx[cur], T.nex[cur][t], T.nex[cur][t^1], l, r)
			os.Exit(0)
		}
		cur = T.nex[cur][t]
		if t != 0 {
			x ^= 1 << i
		}
	}
	return x
}

type pair struct {
	x, y int
}

func query(l, r int) int {
	if l >= r {
		fmt.Println("OOPS")
		os.Exit(0)
	}
	todo := make([]pair, 0)
	cur := l
	for i := 0; i < 31; i++ {
		todo = append(todo, pair{fst[l][i], i})
	}
	todo = append(todo, pair{r, -MOD})
	sort.Slice(todo, func(i, j int) bool {
		if todo[i].x == todo[j].x {
			return todo[i].y < todo[j].y
		}
		return todo[i].x < todo[j].x
	})
	mask := (1 << 31) - 1
	ans := -INF
	for _, t := range todo {
		if cur < t.x {
			bes := T.test(pre[r+1], cur+1, t.x)
			ans = max(ans, mask-bes)
			cur = t.x
		}
		if t.y == -MOD {
			return ans
		}
		mask ^= 1 << t.y
	}
	os.Exit(0)
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
