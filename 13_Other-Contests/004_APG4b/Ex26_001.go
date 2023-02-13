package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	ints := make([]int, 26)
	vecs := make([][]int, 26)

	getInt := func() int {
		var c string
		fmt.Scan(&c)
		if c[0] < 'a' {
			return int(c[0] - '0')
		}
		return ints[int(c[0]-'a')]
	}

	getIntFormula := func() int {
		ret := getInt()
		var op string
		fmt.Scan(&op)
		for op[0] != ';' {
			if op[0] != ' ' {
				i := getInt()
				if op[0] == '+' {
					ret += i
				} else {
					ret -= i
				}
			}
			fmt.Scan(&op)
		}
		return ret
	}

	getVec := func() []int {
		var b string
		fmt.Scan(&b)
		if b[0] >= 'a' {
			ret := make([]int, len(vecs[b[0]-'a']))
			copy(ret, vecs[b[0]-'a'])
			return ret
		}
		ret := make([]int, 0)
		ret = append(ret, getInt())
		var op string
		fmt.Scan(&op)
		for op[0] != ']' {
			ret = append(ret, getInt())
			fmt.Scan(&op)
		}
		return ret
	}

	getVecFormula := func() []int {
		ret := getVec()
		var op string
		fmt.Scan(&op)
		for op[0] != ';' {
			v := getVec()
			for i := 0; i < len(ret); i++ {
				if op[0] == '+' {
					ret[i] += v[i]
				} else {
					ret[i] -= v[i]
				}
			}
			fmt.Scan(&op)
		}
		return ret
	}

	var n int
	fmt.Scan(&n)
	for n > 0 {
		n--
		var com string
		fmt.Scan(&com)
		switch com {
		case "int":
			var id, eq string
			fmt.Scan(&id, &eq)
			ints[int(id[0]-'a')] = getIntFormula()
		case "print_int":
			fmt.Println(getIntFormula())
		case "vec":
			var id, eq string
			fmt.Scan(&id, &eq)
			tmp := getVecFormula()
			vecs[id[0]-'a'] = tmp
		case "print_vec":
			printVec(getVecFormula())
		}
	}
}

func printVec(vec []int) {
	fmt.Print("[ ")
	for i := 0; i < len(vec); i++ {
		fmt.Print(vec[i], " ")
	}
	fmt.Println("]")
}
