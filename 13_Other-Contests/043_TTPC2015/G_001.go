package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)
	n := len(c) / 6
	var at, ai, ae, ac, ah int
	for _, ch := range c {
		switch ch {
		case 't':
			at++
		case 'i':
			ai++
		case 'e':
			ae++
		case 'c':
			ac++
		case 'h':
			ah++
		default:
			fmt.Println("No")
			return
		}

		if at < ai || at < ae*2 || ae < ac || ac < ah || ai < at-n {
			fmt.Println("No")
			return
		}
	}

	if at == n*2 && ai == n && ae == n && ac == n && ah == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
