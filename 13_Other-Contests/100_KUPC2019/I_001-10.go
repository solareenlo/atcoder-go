package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type P struct {
	x bool
	y []int
}

const mushi = 4

var cmd []int
var H [150][]string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var S string
	fmt.Fscan(in, &S)
	if S == "encode" {
		var X int
		fmt.Fscan(in, &X)
		XX := X
		for XX != 0 {
			cmd = append(cmd, XX&1)
			XX >>= 1
		}
		cmd = reverseOrderInt(cmd)
		for i := 0; i < 100; i++ {
			if len(cmd) <= i*mushi+3 {
				break
			}
			insert(&cmd, i*mushi+3, 2)
		}
		for i := range H {
			var t string
			fmt.Fscan(in, &t)
			H[i] = strings.Split(t, "")
		}
		for si := 0; si < 50; si++ {
			for sj := 0; sj < 50; sj++ {
				if H[si][sj] == "#" {
					continue
				}
				if Dfs(0, si, sj) == true {
					nya := efs(0, si, sj)
					nya.y = reverseOrderInt(nya.y)
					ans := make([][]string, 150)
					for i := range ans {
						ans[i] = strings.Split(strings.Repeat("#", 150), "")
					}
					ans[si][sj] = "S"
					for _, x := range nya.y {
						if x == 0 {
							sj++
						} else {
							si++
						}
						ans[si][sj] = "."
					}
					for _, x := range ans {
						fmt.Fprintln(out, strings.Join(x, ""))
						out.Flush()
					}
					return
				}
			}
		}
	} else {
		for i := range H {
			H[i] = strings.Split(strings.Repeat("#", 150), "")
		}
		H[0][0] = "S"
		lastx, lasty := 0, 0
		var send func(int, bool)
		send = func(a int, next bool) {
			if next {
				if a == 0 {
					fmt.Fprintln(out, "? R")
					out.Flush()
				} else {
					fmt.Fprintln(out, "? D")
					out.Flush()
				}
			} else {
				if a == 0 {
					fmt.Fprintln(out, "? L")
					out.Flush()
				} else {
					fmt.Fprintln(out, "? U")
					out.Flush()
				}
			}
		}

		var rec string
		for {
			cur := 0
			send(cur, true)
			fmt.Fscan(in, &rec)
			for i := len(cmd) - 1; i >= 0; i-- {
				send(cmd[i], false)
				fmt.Fscan(in, &rec)
			}
			if rec == "." {
				lasty++
				H[lastx][lasty] = "."
				for i := len(cmd) - 1; i >= 0; i-- {
					if cmd[i] == 1 && H[lastx-1][lasty] == "." {
						lastx--
					}
					if cmd[i] == 0 && H[lastx][lasty-1] == "." {
						lasty--
					}
				}
				for {
					if H[lastx][lasty+1] != "#" {
						send(0, true)
						fmt.Fscan(in, &rec)
						lasty++
						continue
					}
					if H[lastx+1][lasty] != "#" {
						send(1, true)
						fmt.Fscan(in, &rec)
						lastx++
						continue
					}
					break
				}
				cmd = append(cmd, cur)
				continue
			}
			for i := 0; i < len(cmd); i++ {
				send(cmd[i], true)
				fmt.Fscan(in, &rec)
			}
			cur = 1
			send(cur, true)
			fmt.Fscan(in, &rec)
			for i := len(cmd) - 1; i >= 0; i-- {
				send(cmd[i], false)
				fmt.Fscan(in, &rec)
			}
			if rec == "." {
				lastx++
				H[lastx][lasty] = "."
				for i := len(cmd) - 1; i >= 0; i-- {
					if cmd[i] == 1 && H[lastx-1][lasty] != "#" {
						lastx--
					}
					if cmd[i] == 0 && H[lastx][lasty-1] != "#" {
						lasty--
					}
				}
				for {
					if H[lastx][lasty+1] == "." {
						send(0, true)
						fmt.Fscan(in, &rec)
						lasty++
						continue
					}
					if H[lastx+1][lasty] == "." {
						send(1, true)
						fmt.Fscan(in, &rec)
						lastx++
						continue
					}
					break
				}
				cmd = append(cmd, cur)
				continue
			}
			break
		}
		ans := 0
		for i := 0; i < len(cmd); i++ {
			if i%mushi == 3 {
				continue
			}
			ans *= 2
			ans += cmd[i]
		}
		fmt.Fprintln(out, "!", ans)
		out.Flush()
		return
	}
}

func Dfs(depth, x, y int) bool {
	if H[x][y] == "#" {
		return false
	}
	if depth == len(cmd) {
		return true
	}
	if cmd[depth] != 1 { // 横
		if Dfs(depth+1, x, y+1) {
			return true
		}
	}
	if cmd[depth] != 0 { // 縦
		if Dfs(depth+1, x+1, y) {
			return true
		}
	}
	return false
}

func efs(depth, x, y int) P {
	if H[x][y] == "#" {
		return P{false, []int{}}
	}
	if depth == len(cmd) {
		return P{true, []int{}}
	}
	if cmd[depth] != 1 { // 横
		cur := efs(depth+1, x, y+1)
		if cur.x {
			cur.y = append(cur.y, 0)
			return cur
		}
	}
	if cmd[depth] != 0 { // 縦
		cur := efs(depth+1, x+1, y)
		if cur.x {
			cur.y = append(cur.y, 1)
			return cur
		}
	}
	return P{false, []int{}}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
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
