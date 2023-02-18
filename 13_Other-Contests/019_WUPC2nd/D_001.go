package main

import "fmt"

func main() {
	var _1, _2, _3, _4, _5 int
	fmt.Scan(&_1, &_2, &_3, &_4, &_5)
	ans := _5 + _4 + _3
	_1 -= (61*_4 + _3*42)
	_2 -= _3 * 7
	if _2 < 0 {
		_1 += _2 * 8
		_2 = 0
	}
	if _2 > 0 {
		ans += _2 / 8
		if _2%8 != 0 {
			ans++
		}
		_1 -= 61 * (_2 / 8)
		_2 %= 8
		if _2 != 0 {
			_1 -= 125 - _2*8
		} else {
			_1 = 0
		}
	}
	if _1 > 0 {
		ans += _1 / 125
		if _1%125 != 0 {
			ans++
		}
	}
	fmt.Println(ans)
}
