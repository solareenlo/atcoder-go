package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N >= 17 {
		large(N)
	} else {
		small(N)
	}
}

func large(N int) {
	a := 0
	ans := N
	for i := 1; i < N; i++ {
		tmp := max(N-3*i-1, 2*i)
		if ans > tmp {
			a = i
			ans = tmp
		}
	}
	op := make([][][]int, 4)
	for i := range op {
		op[i] = make([][]int, ans)
		for j := range op[i] {
			op[i][j] = make([]int, 3)
			for k := range op[i][j] {
				op[i][j][k] = 1670
			}
		}
	}

	for i := 2; i < a+1; i++ {
		op[0][i-2] = []int{i - 1, i, i}
	}
	op[0][a-1] = []int{a, 2 * a, 2 * a}
	op[0][a] = []int{2 * a, 3*a + 1, 3*a + 1}
	op[0][a+1] = []int{3*a + 1, 4*a + 3, 4*a + 3}
	for i := 4*a + 4; i <= N; i++ {
		op[0][i-(3*a+2)] = []int{i - 1, i, i}
	}

	for i := a + 2; i <= 2*a; i++ {
		op[1][i-(a+2)] = []int{i - 1, i, i}
	}
	for i := a + 1; i < 2*a; i++ {
		op[1][i-2] = []int{a, i, i}
	}
	op[1][2*a-2] = []int{2 * a, 3 * a, 3 * a}
	op[1][2*a-1] = []int{3*a + 1, 4*a + 1, 4*a + 1}

	for i := 2*a + 2; i <= 3*a+1; i++ {
		op[2][i-(2*a+2)] = []int{i - 1, i, i}
	}
	for i := 2*a + 1; i <= 3*a-1; i++ {
		op[2][i-(a+1)] = []int{2 * a, i, i}
	}
	op[2][2*a-1] = []int{3*a + 1, 4*a + 2, 4*a + 2}

	for i := 3*a + 3; i <= 4*a+3; i++ {
		op[3][i-(3*a+3)] = []int{i - 1, i, i}
	}
	for i := 3*a + 2; i <= 4*a; i++ {
		op[3][i-(2*a+1)] = []int{3*a + 1, i, i}
	}

	if N%5 == 3 {
		ans--
		op[0][ans-1] = []int{N - 2, N, N}
		op[1][ans-1] = op[1][ans-2]
		op[2][ans-1] = op[2][ans-2]
		op[1][ans-2] = []int{N - 1, N, N}
		op[2][ans-2] = []int{N - 2, N - 1, N - 1}
		op[3][ans-1] = []int{N - 3, N - 1, N - 1}
	}
	// check
	type pair struct {
		x, y int
	}
	p := make([]pair, N+1)
	for i := 0; i < N+1; i++ {
		p[i] = pair{i, i + 1}
	}
	for i := 0; i < ans; i++ {
		tmp := make([]pair, 4)
		for k := 0; k < 4; k++ {
			if op[k][i][0] == 1670 {
				continue
			}
			tmp[k] = pair{p[op[k][i][0]].x, p[op[k][i][1]].y}
		}
		for k := 0; k < 4; k++ {
			if op[k][i][0] == 1670 {
				continue
			}
			p[op[k][i][2]] = tmp[k]
		}
	}

	// output
	fmt.Println(ans)
	for i := 0; i < ans; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 3; k++ {
				fmt.Print(op[j][i][(k+2)%3])
				if k == 2 {
					fmt.Println()
				} else {
					fmt.Print(" ")
				}
			}
		}
	}
}

func small(N int) {
	po := "167 167 167\n"
	if N == 2 {
		fmt.Println(N - 1)
		for i := 1; i < N; i++ {
			fmt.Println(i+1, i, i+1)
			for k := 1; k < 4; k++ {
				fmt.Print(po)
			}
		}
	} else if N == 3 || N == 4 {
		fmt.Println(2)
		for i := 0; i < 4; i++ {
			fmt.Println(i+2, i+1, i+2)
		}
		fmt.Println(3, 1, 3)
		fmt.Println(4, 2, 4)
		fmt.Print(po)
		fmt.Print(po)
	} else if N <= 8 {
		fmt.Println(3)
		for i := 0; i < 4; i++ {
			fmt.Println(i*2+2, i*2+1, i*2+2)
		}
		for i := 0; i < 4; i++ {
			tmp := 2 + 4*(i/2)
			fmt.Println(tmp+1+i%2, tmp, tmp+1+i%2)
		}
		for i := 0; i < 4; i++ {
			fmt.Println(5+i, 4, 5+i)
		}
	} else if N <= 11 {
		fmt.Println(4)
		for i := 0; i < 4; i++ {
			fmt.Println(i*2+2, i*2+1, i*2+2)
		}
		fmt.Print("3 2 3\n4 2 4\n8 6 8\n10 9 10\n")
		fmt.Print("5 4 5\n6 4 6\n8 4 8\n11 10 11\n")
		fmt.Print("7 6 7\n9 8 9\n10 8 10\n11 8 11\n")
	} else if N <= 13 {
		fmt.Println(5)
		for i := 0; i < 4; i++ {
			fmt.Println(i*2+2, i*2+1, i*2+2)
		}
		fmt.Print("3 2 3\n4 2 4\n8 6 8\n10 9 10\n")
		fmt.Print("5 4 5\n6 4 6\n8 4 8\n11 10 11\n")
		fmt.Print("7 6 7\n9 8 9\n13 12 13\n11 8 11\n")
		fmt.Print("12 11 12\n13 11 13\n10 8 10\n")
		fmt.Print(po)
	} else if N <= 16 {
		fmt.Println(6)
		for i := 0; i < 4; i++ {
			fmt.Println(i*2+2, i*2+1, i*2+2)
		}
		fmt.Print("3 2 3\n4 2 4\n8 6 8\n10 9 10\n")
		fmt.Print("5 4 5\n6 4 6\n8 4 8\n11 10 11\n")
		fmt.Print("15 14 15\n9 8 9\n13 12 13\n11 8 11\n")
		fmt.Print("12 11 12\n13 11 13\n10 8 10\n16 15 16\n")
		fmt.Print("14 13 14\n15 13 15\n16 13 16\n7 6 7\n")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
