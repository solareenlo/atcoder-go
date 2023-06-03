package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var S [1001]string
	var T [1001]int

	var N int
	fmt.Fscan(in, &N)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &S[i], &T[i])
	}
	ans := make([]string, 0)
	for i := 0; i <= 9999; i++ {
		ID := strconv.Itoa(i)
		for len(ID) < 4 {
			ID = "0" + ID
		}
		flag := true
		for j := 1; j <= N; j++ {
			if Hantei(S[j], ID) != T[j] {
				flag = false
			}
		}
		if flag == true {
			ans = append(ans, ID)
		}
	}
	if len(ans) != 1 {
		fmt.Println("Can't Solve")
	} else {
		fmt.Println(ans[0])
	}
}

func Hantei(A1, A2 string) int {
	diff := 0
	for i := 0; i < 4; i++ {
		if A1[i] != A2[i] {
			diff += 1
		}
	}
	if diff == 0 {
		return 1
	}
	if diff == 1 {
		return 2
	}
	return 3
}
