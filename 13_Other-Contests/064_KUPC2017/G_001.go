package main

import (
	"fmt"
	"os"
)

var x int64

func encode() {
	ofs, err := os.Create("/tmp/hoge")
	if err != nil {
		panic(err)
	}
	defer ofs.Close()
	_, err = fmt.Fprint(ofs, x)
	if err != nil {
		panic(err)
	}

}

func decode() {
	ifs, err := os.Open("/tmp/hoge")
	if err != nil {
		panic(err)
	}
	defer ifs.Close()
	_, err = fmt.Fscanln(ifs, &x)
	if err != nil {
		panic(err)
	}

}

func main() {
	var s string
	fmt.Scanln(&s)
	if s == "encode" {
		var n, m, xx, y int
		fmt.Scanln(&n, &m)
		for i := 1; i <= m; i++ {
			fmt.Scanln(&xx, &y)
		}
		fmt.Scanln(&x)
		fmt.Println("0")
		encode()
	} else {
		decode()
		fmt.Println(x)
	}
}
