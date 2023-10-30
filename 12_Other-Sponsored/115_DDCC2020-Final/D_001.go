package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 4557430888798830399
const MXN = 800008

var cnt, tot, maxu, dead, now, now2, all int
var to, v, nex, head, sum, ide, ans, pos [MXN]int
var loop, val, id []int
var f [MXN][3]int
var vis [MXN]bool
var in [MXN][3]int

func add(x, y, z int) {
	cnt++
	to[cnt] = y
	v[cnt] = z
	nex[cnt] = head[x]
	head[x] = cnt
}

var st []int

type Node struct {
	maxn, max1, max2, l, l1, r, r2 int
}
type NODE struct {
	l, r int
	nd   Node
}

var node [3400001]NODE

func plus(nd1, nd2 Node) Node {
	var nd Node
	nd.l = nd1.l1
	nd.r = nd2.r2
	nd.maxn = nd1.max1 + nd2.max2
	if nd1.maxn > nd.maxn {
		nd.maxn = nd1.maxn
		nd.l = nd1.l
		nd.r = nd1.r
	}
	if nd2.maxn > nd.maxn {
		nd.maxn = nd2.maxn
		nd.l = nd2.l
		nd.r = nd2.r
	}
	if nd1.max1 > nd2.max1 {
		nd.l1 = nd1.l1
	} else {
		nd.l1 = nd2.l1
	}
	nd.max1 = Max(nd1.max1, nd2.max1)
	if nd1.max2 > nd2.max2 {
		nd.r2 = nd1.r2
	} else {
		nd.r2 = nd2.r2
	}
	nd.max2 = Max(nd1.max2, nd2.max2)
	return nd
}

func Build(p, l, r int) {
	node[p].l = l
	node[p].r = r
	if l == r {
		node[p].nd.l = l
		node[p].nd.l1 = l
		node[p].nd.r = l
		node[p].nd.r2 = l
		val[l] += val[l-1]
		node[p].nd.max2 = f[loop[l]][0] + val[l-1]
		node[p].nd.max1 = f[loop[l]][0] - val[l-1]
		return
	}
	mid := l + ((r - l) >> 1)
	Build(p*2, l, mid)
	Build(p*2+1, mid+1, r)
	node[p].nd = plus(node[p*2].nd, node[p*2+1].nd)
}

func change(p, x, y int) {
	if node[p].l == node[p].r {
		node[p].nd.max2 = y + val[x]
		node[p].nd.max1 = y - val[x]
		return
	}
	if x <= node[p*2].r {
		change(p*2, x, y)
	} else {
		change(p*2+1, x, y)
	}
	node[p].nd = plus(node[p*2].nd, node[p*2+1].nd)
}

func ask(p, l, r int) Node {
	if l <= node[p].l && r >= node[p].r {
		return node[p].nd
	}
	if l > node[p*2].r {
		return ask(p*2+1, l, r)
	}
	if r <= node[p*2].r {
		return ask(p*2, l, r)
	}
	return plus(ask(p*2, l, r), ask(p*2+1, l, r))
}

func findloop(k, from int) {
	st = append(st, k)
	for i := head[k]; i > 0; i = nex[i] {
		if i == (from ^ 1) {
			continue
		}
		if sum[to[i]] != 0 {
			var y int
			sum[k] = v[i]
			ide[k] = i
			for {
				tot++
				y = st[len(st)-1]
				loop[tot] = y
				vis[loop[tot]] = true
				id[tot] = ide[y]
				val[tot] = sum[y]
				st = st[:len(st)-1]
				if y == to[i] {
					break
				}
			}
		} else {
			sum[k] = v[i]
			ide[k] = i
			findloop(to[i], i)
		}
		if tot != 0 {
			return
		}
	}
	st = st[:len(st)-1]
}

func dfs(k, from int) {
	for i := head[k]; i > 0; i = nex[i] {
		if i == (from^1) || vis[to[i]] {
			continue
		}
		dfs(to[i], i)
		tmp := Max(in[to[i]][0], f[to[i]][0]+f[to[i]][1])
		if tmp >= in[k][0] {
			in[k][1] = in[k][0]
			in[k][0] = tmp
		} else if tmp >= in[k][1] {
			in[k][1] = tmp
		}
		if f[to[i]][0]+v[i] >= f[k][0] {
			f[k][2] = f[k][1]
			f[k][1] = f[k][0]
			f[k][0] = f[to[i]][0] + v[i]
		} else if f[to[i]][0]+v[i] >= f[k][1] {
			f[k][2] = f[k][1]
			f[k][1] = f[to[i]][0] + v[i]
		} else if f[to[i]][0]+v[i] >= f[k][2] {
			f[k][2] = f[to[i]][0] + v[i]
		}
	}
}

func dfs2(k, maxn, from, up int) {
	pre := maxu
	ppr := dead
	for i, upon := head[k], 0; i > 0; i = nex[i] {
		if vis[to[i]] || i == (from^1) {
			continue
		}
		if in[to[i]][0] == in[k][0] || f[to[i]][0]+f[to[i]][1] == in[k][0] {
			dead = Max(dead, in[k][1])
		} else {
			dead = Max(dead, in[k][0])
		}
		if f[to[i]][0]+v[i] == f[k][0] {
			dead = Max(dead, f[k][1]+maxn, f[k][1]+f[k][2])
		} else if f[to[i]][0]+v[i] == f[k][1] {
			dead = Max(dead, f[k][0]+maxn, f[k][0]+f[k][2])
		} else {
			dead = Max(dead, f[k][0]+f[k][1], maxn+f[k][0])
		}
		if f[to[i]][0]+v[i] == f[k][0] {
			upon = f[k][1]
			maxu = Max(maxu, up+f[k][1])
			dfs2(to[i], Max(maxn, f[k][1])+v[i], i, up+v[i])
		} else {
			upon = f[k][0]
			maxu = Max(maxu, up+f[k][0])
			dfs2(to[i], Max(maxn, f[k][0])+v[i], i, up+v[i])
		}
		ans[pos[i]] = Max(dead, in[to[i]][0], f[to[i]][0]+f[to[i]][1], now+Max(maxu, upon+up), now2)
		if f[to[i]][0]+v[i] == f[k][0] {
			ans[pos[i]] = Max(ans[pos[i]], f[k][1]+f[k][2], f[k][1]+maxn)
		} else if f[to[i]][0]+v[i] == f[k][1] {
			ans[pos[i]] = Max(ans[pos[i]], f[k][0]+f[k][2], f[k][0]+maxn)
		} else {
			ans[pos[i]] = Max(ans[pos[i]], f[k][0]+f[k][1], f[k][0]+maxn)
		}
		maxu = pre
		dead = ppr
	}
}

func calc2() Node {
	var nd1, nd2 Node
	nd1.maxn = 0
	for i, j := 2, 1; i <= tot*2; i++ {
		for (val[i]-val[j])*2 > all {
			j++
		}
		nd2 = ask(1, j, i)
		if nd2.maxn > nd1.maxn {
			nd1 = nd2
		}
	}
	return nd1
}

func calc(ex int) {
	if ex > tot {
		ex -= tot
	}
	change(1, ex, 0)
	change(1, ex+tot, 0)
	dead = 0
	now = -INF
	now2 = calc2().maxn
	for i := 1; i <= tot; i++ {
		if i != ex {
			now2 = Max(now2, in[loop[i]][0], f[loop[i]][0]+f[loop[i]][1])
		}
	}
	for i := ex + tot - 1; i > ex; i-- {
		if (val[i]-val[ex])*2 <= val[tot+1] {
			now = Max(now, val[i]+f[loop[i]][0]-val[ex])
		}
	}
	for i := ex + 1; i < ex+tot; i++ {
		if (val[ex+tot]-val[i])*2 <= val[tot+1] {
			now = Max(now, f[loop[i]][0]-val[i]+val[ex+tot])
		}
	}
	dfs2(loop[ex], 0, 0, 0)
	change(1, ex, f[loop[ex]][0])
	change(1, ex+tot, f[loop[ex]][0])
}

func main() {
	IN := bufio.NewReader(os.Stdin)

	cnt = 1
	loop = make([]int, MXN)
	val = make([]int, MXN)
	id = make([]int, MXN)

	var n int
	fmt.Fscan(IN, &n)
	pos2 := 0
	var nd Node
	for i := 1; i <= n; i++ {
		var x, y, z int
		fmt.Fscan(IN, &x, &y, &z)
		add(x, y, z)
		add(y, x, z)
		pos[cnt] = i
		pos[cnt-1] = i
	}
	findloop(1, 0)
	reverseOrderInt(loop[1 : tot+1])
	reverseOrderInt(val[1 : tot+1])
	reverseOrderInt(id[1 : tot+1])
	id[0] = id[tot]
	var max2 int
	for i := 1; i <= tot; i++ {
		loop[i+tot] = loop[i]
		val[i+tot] = val[i]
		now = i
		dfs(loop[i], 0)
		max2 = Max(max2, in[loop[i]][0], f[loop[i]][0]+f[loop[i]][1])
		if Max(in[loop[i]][0], f[loop[i]][0]+f[loop[i]][1]) == max2 {
			pos2 = i
		}
	}
	Build(1, 1, tot*2)
	all = val[tot]
	for i := tot * 2; i >= 1; i-- {
		val[i] = val[i-1]
	}
	nd = calc2()
	for i := 1; i <= tot; i++ {
		ans[pos[id[i-1]]] = Max(max2, ask(1, i, i+tot-1).maxn)
	}
	if nd.maxn <= max2 {
		calc(pos2)
	} else {
		calc(nd.l)
		calc(nd.r)
	}
	for i := 1; i <= n; i++ {
		if ans[i] == 0 {
			fmt.Println(Max(nd.maxn, max2))
		} else {
			fmt.Println(ans[i])
		}
	}
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func Max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
