package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const BASE = 8
const KETA = 20

var in = bufio.NewReader(os.Stdin)
var Out = bufio.NewWriter(os.Stdout)

func main() {
	defer Out.Flush()
	var tag string
	fmt.Fscan(in, &tag)

	if tag == "encode" {
		encode()
	} else {
		decode()
	}
}

func encode() {
	var xxx int
	fmt.Fscan(in, &xxx)
	const w = 150

	bd := make([][]string, w)
	for i := 0; i < w; i++ {
		var s string
		fmt.Fscan(in, &s)
		bd[i] = strings.Split(s, "")
	}

	type pair struct {
		x, y int
	}

	var pos pair
	flag := false
	for y := 0; y < len(bd); y++ {
		for x := 11; x < len(bd[y])-11; x++ {
			var found func() bool
			found = func() bool {
				xx := xxx
				rest := KETA
				for cy := y; ; cy++ {
					if bd[cy][x] == "#" {
						return false
					}
					llen := 0
					for llen < BASE && bd[cy][x-(llen+1)] == "." {
						llen += 1
					}
					if xx%BASE+1 <= llen {
						xx /= BASE
						rest -= 1
						if rest == 0 {
							return true
						}
					}
					rlen := 0
					for rlen < BASE && bd[cy][x+(rlen+1)] == "." {
						rlen += 1
					}
					if xx%BASE+1 <= rlen {
						xx /= BASE
						rest -= 1
						if rest == 0 {
							return true
						}
					}
				}
				return false
			}
			if found() {
				pos = pair{x, y}
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}

	var out [150][]string
	for i := range out {
		out[i] = make([]string, 150)
		for j := range out[i] {
			out[i][j] = "#"
		}
	}

	xx := xxx
	dir := -1
	cy := pos.y
	rest := KETA
	for rest > 0 {
		out[cy][pos.x] = "."
		llen := 0
		for llen < BASE && pos.x+(llen+1)*dir >= 0 && bd[cy][pos.x+(llen+1)*dir] == "." {
			llen++
		}

		if xx%BASE+1 <= llen {
			code := xx%BASE + 1
			for ii := 0; ii < code; ii++ {
				if pos.x+(ii+1)*dir >= 0 {
					out[cy][pos.x+(ii+1)*dir] = "."
				}
			}

			xx /= BASE
			rest--
		}
		if dir < 0 {
			dir = 1
		} else {
			dir = -1
			cy++
		}
	}

	out[pos.y][pos.x] = "S"

	for _, row := range out {
		fmt.Fprintln(Out, strings.Join(row, ""))
		Out.Flush()
	}
}

func decode() {

	rest := KETA
	dir := -1
	cy := 0
	ans := 0
	keta := 1

	for rest > 0 {
		width := 0
		for cx := 1; ; cx++ {
			for i := 0; i < cy; i++ {
				mov("D")
			}
			ru := false
			for i := 0; i < cx; i++ {
				if dir > 0 {
					ru = mov("R")
				} else {
					ru = mov("L")
				}
			}
			lu := false
			for i := 0; i < cx-1; i++ {
				if dir > 0 {
					lu = mov("L")
				} else {
					lu = mov("R")
				}
			}
			if cy > 0 {
				for i := 0; i < cy-1; i++ {
					mov("U")
				}
				if mov("U") {
					break
				}
			} else {
				if cx == 1 {
					if ru {
						break
					}
				} else {
					if lu {
						break
					}
				}
			}
			if dir > 0 {
				lu = mov("L")
			} else {
				lu = mov("R")
			}
			if cy > 0 {
				for !mov("U") {
				}
			}
			width = cx
		}
		if dir < 0 {
			dir = 1
		} else {
			dir = -1
			cy++
		}
		if width == 0 {
			continue
		}
		ans += (width - 1) * keta
		keta *= BASE
		rest--
	}
	fmt.Fprintln(Out, "!", ans)
	Out.Flush()
}

func mov(dir string) bool {
	fmt.Fprintln(Out, "?", dir)
	Out.Flush()
	var s string
	fmt.Fscan(in, &s)
	return s == "S"
}
