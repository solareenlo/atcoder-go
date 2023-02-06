package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &s[i])
	}
	s = transpose(s)
	sort.Strings(s)

	t := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &t[i])
	}
	t = transpose(t)
	sort.Strings(t)

	if reflect.DeepEqual(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func transpose(slice []string) []string {
	xl := len(slice[0])
	yl := len(slice)
	s := make([][]string, yl)
	for i := 0; i < yl; i++ {
		s[i] = strings.Split(slice[i], "")
	}
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = s[j][i]
		}
	}
	res := make([]string, xl)
	for i := 0; i < xl; i++ {
		res[i] = strings.Join(result[i], "")
	}
	return res
}
