package main

import "fmt"

const N = 6
const A = 4
const B = 4

func main() {
	cnf := make([][]int, 0)
	t := make([]int, 0)

	for a := 0; a < A; a++ {
		for b := 0; b < B; b++ {
			for n := 0; n < N*N; n++ {
				t = append(t, ctoi(n, a, b))
			}
			cnf = append(cnf, t)
			t = make([]int, 0)
		}
	}

	for a := 0; a < A; a++ {
		for b := 0; b < B; b++ {
			for n1 := 0; n1 < N*N; n1++ {
				for n2 := 0; n2 < N*N; n2++ {
					if n1 < n2 {
						t = append(t, -ctoi(n1, a, b))
						t = append(t, -ctoi(n2, a, b))
						cnf = append(cnf, t)
						t = make([]int, 0)
					}
				}
			}
		}
	}

	for n := 0; n < N*N; n++ {
		for m1 := 0; m1 < A*B; m1++ {
			for m2 := 0; m2 < A*B; m2++ {
				if m1 < m2 {
					t = append(t, -ctoi(n, m1/B, m1%B))
					t = append(t, -ctoi(n, m2/B, m2%B))
					cnf = append(cnf, t)
					t = make([]int, 0)
				}
			}
		}
	}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			for a := 0; a < A; a++ {
				for b2 := 0; b2 < B-1; b2++ {
					t = append(t, -ctoi(y*N+x, a, b2+1))
					for i := 0; i < 4; i++ {
						if 0 <= y+dy[i] && y+dy[i] < N && 0 <= x+dx[i] && x+dx[i] < N {
							for b1 := 0; b1 < b2+1; b1++ {
								t = append(t, ctoi((y+dy[i])*N+(x+dx[i]), a, b1))
							}
						}
					}
					cnf = append(cnf, t)
					t = make([]int, 0)
				}
			}
		}
	}

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			for a2 := 0; a2 < A; a2++ {
				for b2 := 0; b2 < B; b2++ {
					for a1 := 0; a1 < a2; a1++ {
						for b1 := 0; b1 < B; b1++ {
							for i := 0; i < 4; i++ {
								if 0 <= y+dy[i] && y+dy[i] < N && 0 <= x+dx[i] && x+dx[i] < N {
									t = append(t, -ctoi(y*N+x, a1, b1))
									t = append(t, -ctoi((y+dy[i])*N+(x+dx[i]), a2, b2))
									cnf = append(cnf, t)
									t = make([]int, 0)
								}
							}
						}
					}
				}
			}
		}
	}

	for n := 0; n < N*N; n++ {
		t = append(t, n+1)
		for a := 0; a < A; a++ {
			for b := 0; b < B; b++ {
				t = append(t, ctoi(n, a, b))
			}
		}
		cnf = append(cnf, t)
		t = make([]int, 0)
	}
	for n := 0; n < N*N; n++ {
		for a := 0; a < A; a++ {
			for b := 0; b < B; b++ {
				t = append(t, -(n + 1))
				t = append(t, -ctoi(n, a, b))
				cnf = append(cnf, t)
				t = make([]int, 0)
			}
		}
	}

	for _, i := range cnf {
		for _, j := range i {
			fmt.Printf("%d ", j)
		}
		fmt.Println(0)
	}
	fmt.Println(0)
}

func ctoi(i, a, b int) int {
	return i + (a*B+b+1)*N*N + 1
}
