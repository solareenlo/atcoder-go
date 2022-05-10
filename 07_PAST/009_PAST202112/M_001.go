package main

import "fmt"

type data struct{ key, id int }

const mask = 65535

var (
	y = [100001]data{}
	z = [100001]int{}
)

func radix_sort(n int, x []data) {
	l := make([]int, 1<<16)
	m := make([]int, 1<<16)
	for j, k := 0, 0; j < 4; j, k = j+1, k+16 {
		for i := 0; i <= mask; i++ {
			m[i] = 0
		}
		for i := 0; i < n; i++ {
			z[i] = (x[i].key >> k) & mask
			m[z[i]]++
		}
		if m[0] == n {
			break
		}
		l[0] = 0
		for i := 0; i < mask; i++ {
			l[i+1] = l[i] + m[i]
		}
		for i := 0; i < n; i++ {
			y[l[z[i]]] = x[i]
			l[z[i]]++
		}
		for i := 0; i < n; i++ {
			x[i] = y[i]
		}
	}
}

func encode(S string, id int) int {
	var i int
	ans := 0
	for i := range S {
		ans = (ans << 5) + int(S[i]-'a') + 1
	}
	for ; i < 5; i++ {
		ans <<= 5
	}
	return (ans << 20) + id
}

type list struct {
	id    int
	size  int
	val   int
	child [2]*list
	par   *list
}

func add_node(root, p *list) {
	root.size++
	if p.val < root.val {
		if root.child[0] == nil {
			root.child[0] = p
			p.par = root
			p.child[0] = nil
			p.child[1] = nil
			p.size = 1
		} else {
			add_node(root.child[0], p)
		}
	} else {
		if root.child[1] == nil {
			root.child[1] = p
			p.par = root
			p.child[0] = nil
			p.child[1] = nil
			p.size = 1
		} else {
			add_node(root.child[1], p)
		}
	}
}

func delete_node(p *list) {
	for ; p.par != nil; p = p.par {
		p.size--
	}
	p.size--
}

func search(root *list, k int) *list {
	l := root.child[0].size
	if root.child[0] == nil {
		l = 0
	}
	r := root.child[1].size
	if root.child[1] == nil {
		r = 0
	}
	if l == k-1 && l+r < root.size {
		return root
	} else if l >= k {
		return search(root.child[0], k)
	} else {
		tmp := 0
		if l+r < root.size {
			tmp = 1
		}
		return search(root.child[1], k-l-tmp)
	}
}

const THR = 5000

var bit = [21]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576}

func main() {
	var N, Q int
	fmt.Scan(&N, &Q)

	S := make([]string, N+1)
	d := make([]data, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&S[i])
		d[i-1].key = encode(S[i], i)
		d[i-1].id = i
	}

	if N == 1 {
		var j int
		var T string
		for i := 1; i <= Q; i++ {
			fmt.Scan(&j, &T)
		}
		fmt.Println(T)
		return
	}

	radix_sort(N, d)

	var k int
	for k = 0; bit[k]-1 < N; k++ {
	}
	k--
	i := bit[k] - 1
	ld := make([]list, 200001)
	ld[i].val = d[i].key
	ld[i].id = d[i].id
	ld[i].size = 1
	ld[i].par = nil
	ld[i].child[0] = nil
	ld[i].child[1] = nil
	root := &(ld[i])
	for k--; k >= 0; k-- {
		for i = bit[k] - 1; i < N; i += bit[k+1] {
			ld[i].val = d[i].key
			ld[i].id = d[i].id
			add_node(root, &(ld[i]))
		}
	}
	for q := 1; q <= Q; q++ {
		var T string
		fmt.Scan(&k, &T)
		p := search(root, k)
		delete_node(p)
		i = p.id
		S[i] = T
		ld[N+q].val = encode(S[i], i)
		ld[N+q].id = i
		add_node(root, &(ld[N+q]))
	}

	for i := 1; i <= N; i++ {
		fmt.Print(S[i], " ")
	}
}
