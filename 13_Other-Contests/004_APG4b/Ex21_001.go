package main

import "fmt"

func f0(N int) int {
	return 1
}

func f1(N, M int) int {
	s := 0
	for i := 0; i < N; i++ {
		s++
	}
	for i := 0; i < M; i++ {
		s++
	}
	return s
}

func f2(N int) int {
	s := 0
	for i := 0; i < N; i++ {
		t := N
		cnt := 0
		for t > 0 {
			cnt++
			t /= 2
		}
		s += cnt
	}
	return s
}

func f3(N int) int {
	s := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s++
		}
	}
	return s
}

func f4(N int) int {
	s := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			s += i + j
		}
	}
	return s
}

func f5(N, M int32) int32 {
	s := int32(0)
	for i := int32(0); i < N; i++ {
		for j := int32(0); j < M; j++ {
			s += i + j
		}
	}
	return s
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	a0, a1, a2, a3, a4, a5 := -1, -1, -1, -1, -1, -1

	a0 = f0(N)
	a1 = f1(N, M)
	a2 = f2(N)
	a3 = f3(N)
	// a4 = f4(N)
	a5 = int(f5(int32(N), int32(M)))

	fmt.Println("f0:", a0)
	fmt.Println("f1:", a1)
	fmt.Println("f2:", a2)
	fmt.Println("f3:", a3)
	fmt.Println("f4:", a4)
	fmt.Println("f5:", a5)
}
