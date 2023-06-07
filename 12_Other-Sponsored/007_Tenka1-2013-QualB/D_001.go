package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const W = 12
	const H = 12

	var b, flag [12][12]int

	var i, j int
	for i = 0; i < H; i++ {
		for j = 0; j < W; j++ {
			fmt.Fscan(in, &b[i][j])
			b[i][j]--
		}
	}

	k := 0
	var x, y, z, head, tail, tmp int
	var q [144][2]int
	var par [12][12][2]int
	var ans [2500][2]int
	for z = 0; z < 123; z++ {
		for b[z/W][z%W] != z {
			for x = 0; x < H; x++ {
				for y = 0; y < W; y++ {
					if b[x][y] == z {
						break
					}
				}
				if y < W {
					break
				}
			}
			if z%W == W-1 && x == z/W+1 && y == z%W && b[x-1][y] != -1 {
				for i = 0; i < H; i++ {
					for j = 0; j < W; j++ {
						par[i][j][0] = -1 - flag[i][j]
					}
				}
				tail = 0
				par[x][y][0] = -2
				par[x+1][y][0] = x
				par[x+1][y][1] = y
				q[tail][0] = x + 1
				q[tail][1] = y
				tail++
				for head = 0; head < tail; head++ {
					i = q[head][0]
					j = q[head][1]
					if b[i][j] == -1 {
						break
					}
					if i > 0 && par[i-1][j][0] == -1 {
						par[i-1][j][0] = i
						par[i-1][j][1] = j
						q[tail][0] = i - 1
						q[tail][1] = j
						tail++
					}
					if i < H-1 && par[i+1][j][0] == -1 {
						par[i+1][j][0] = i
						par[i+1][j][1] = j
						q[tail][0] = i + 1
						q[tail][1] = j
						tail++
					}
					if j > 0 && par[i][j-1][0] == -1 {
						par[i][j-1][0] = i
						par[i][j-1][1] = j
						q[tail][0] = i
						q[tail][1] = j - 1
						tail++
					}
					if j < W-1 && par[i][j+1][0] == -1 {
						par[i][j+1][0] = i
						par[i][j+1][1] = j
						q[tail][0] = i
						q[tail][1] = j + 1
						tail++
					}
				}
				if head == tail {
					return
				}
				for i != x || j != y {
					ans[k][0] = b[par[i][j][0]][par[i][j][1]]
					if par[i][j][0] == i+1 {
						ans[k][1] = 0
						k++
					}
					if par[i][j][0] == i-1 {
						ans[k][1] = 1
						k++
					}
					if par[i][j][1] == j+1 {
						ans[k][1] = 2
						k++
					}
					if par[i][j][1] == j-1 {
						ans[k][1] = 3
						k++
					}
					b[i][j] = b[par[i][j][0]][par[i][j][1]]
					b[par[i][j][0]][par[i][j][1]] = -1
					tmp = par[i][j][0]
					j = par[i][j][1]
					i = tmp
				}
				ans[k][0] = b[x-1][y]
				ans[k][1] = 1
				k++
				b[x][y] = b[x-1][y]
				b[x-1][y] = -1
				for i = 0; i < H; i++ {
					for j = 0; j < W; j++ {
						par[i][j][0] = -1 - flag[i][j]
					}
				}
				tail = 0
				par[x-1][y][0] = -2
				par[x][y][0] = -2
				par[x+1][y][0] = -2
				par[x][y-1][0] = x
				par[x][y-1][1] = y
				q[tail][0] = x
				q[tail][1] = y - 1
				tail++
				for head = 0; head < tail; head++ {
					i = q[head][0]
					j = q[head][1]
					if b[i][j] == -1 {
						break
					}
					if i > 0 && par[i-1][j][0] == -1 {
						par[i-1][j][0] = i
						par[i-1][j][1] = j
						q[tail][0] = i - 1
						q[tail][1] = j
						tail++
					}
					if i < H-1 && par[i+1][j][0] == -1 {
						par[i+1][j][0] = i
						par[i+1][j][1] = j
						q[tail][0] = i + 1
						q[tail][1] = j
						tail++
					}
					if j > 0 && par[i][j-1][0] == -1 {
						par[i][j-1][0] = i
						par[i][j-1][1] = j
						q[tail][0] = i
						q[tail][1] = j - 1
						tail++
					}
					if j < W-1 && par[i][j+1][0] == -1 {
						par[i][j+1][0] = i
						par[i][j+1][1] = j
						q[tail][0] = i
						q[tail][1] = j + 1
						tail++
					}
				}
				if head == tail {
					return
				}
				for i != x || j != y {
					ans[k][0] = b[par[i][j][0]][par[i][j][1]]
					if par[i][j][0] == i+1 {
						ans[k][1] = 0
						k++
					}
					if par[i][j][0] == i-1 {
						ans[k][1] = 1
						k++
					}
					if par[i][j][1] == j+1 {
						ans[k][1] = 2
						k++
					}
					if par[i][j][1] == j-1 {
						ans[k][1] = 3
						k++
					}
					b[i][j] = b[par[i][j][0]][par[i][j][1]]
					b[par[i][j][0]][par[i][j][1]] = -1
					tmp = par[i][j][0]
					j = par[i][j][1]
					i = tmp
				}
				ans[k][0] = z
				ans[k][1] = 0
				k++
				ans[k][0] = z
				ans[k][1] = 0
				k++
				b[x-1][y] = z
				b[x+1][y] = -1
				break
			}
			for i = 0; i < H; i++ {
				for j = 0; j < W; j++ {
					par[i][j][0] = -1 - flag[i][j]
				}
			}
			tail = 0
			par[x][y][0] = -2
			if x > z/W && par[x-1][y][0] == -1 {
				par[x-1][y][0] = x
				par[x-1][y][1] = y
				q[tail][0] = x - 1
				q[tail][1] = y
				tail++
			}
			if y > z%W && par[x][y-1][0] == -1 {
				par[x][y-1][0] = x
				par[x][y-1][1] = y
				q[tail][0] = x
				q[tail][1] = y - 1
				tail++
			}
			if y < z%W && par[x][y+1][0] == -1 {
				par[x][y+1][0] = x
				par[x][y+1][1] = y
				q[tail][0] = x
				q[tail][1] = y + 1
				tail++
			}
			for head = 0; head < tail; head++ {
				i = q[head][0]
				j = q[head][1]
				if b[i][j] == -1 {
					break
				}
				if i > 0 && par[i-1][j][0] == -1 {
					par[i-1][j][0] = i
					par[i-1][j][1] = j
					q[tail][0] = i - 1
					q[tail][1] = j
					tail++
				}
				if i < H-1 && par[i+1][j][0] == -1 {
					par[i+1][j][0] = i
					par[i+1][j][1] = j
					q[tail][0] = i + 1
					q[tail][1] = j
					tail++
				}
				if j > 0 && par[i][j-1][0] == -1 {
					par[i][j-1][0] = i
					par[i][j-1][1] = j
					q[tail][0] = i
					q[tail][1] = j - 1
					tail++
				}
				if j < W-1 && par[i][j+1][0] == -1 {
					par[i][j+1][0] = i
					par[i][j+1][1] = j
					q[tail][0] = i
					q[tail][1] = j + 1
					tail++
				}
			}
			if head == tail {
				return
			}
			for i != x || j != y {
				ans[k][0] = b[par[i][j][0]][par[i][j][1]]
				if par[i][j][0] == i+1 {
					ans[k][1] = 0
					k++
				}
				if par[i][j][0] == i-1 {
					ans[k][1] = 1
					k++
				}
				if par[i][j][1] == j+1 {
					ans[k][1] = 2
					k++
				}
				if par[i][j][1] == j-1 {
					ans[k][1] = 3
					k++
				}
				b[i][j] = b[par[i][j][0]][par[i][j][1]]
				b[par[i][j][0]][par[i][j][1]] = -1
				tmp = par[i][j][0]
				j = par[i][j][1]
				i = tmp
			}
		}
		flag[z/W][z%W] = 1
	}

	fmt.Println(k)
	for i = 0; i < k; i++ {
		fmt.Printf("%d ", ans[i][0]+1)
		if ans[i][1] == 0 {
			fmt.Printf("up\n")
		} else if ans[i][1] == 1 {
			fmt.Printf("down\n")
		} else if ans[i][1] == 2 {
			fmt.Printf("left\n")
		} else {
			fmt.Printf("right\n")
		}
	}
}
