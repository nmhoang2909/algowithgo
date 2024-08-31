package main

import (
	"strings"
)

/*
5 % 2 = 1 => 1
5 / 2 = 2

2 % 2 = 0 => 01
2 / 2 = 1

1 % 2 = 1 => 101
1 / 2 = 0 => stop
*/

/*
charset = [0123456789ABCDEF]
*/
func DecToBase(dec, base int) string {
	// var ret string
	// for dec > 0 {
	// 	rem := dec % base
	// 	ret = fmt.Sprintf("%X%s", rem, ret)
	// 	dec = dec / base
	// }
	// return ret

	charset := "0123456789ABCDEFGH"
	var ret strings.Builder
	for dec > 0 {
		rem := dec % base
		ret.WriteByte(charset[rem])
		dec = dec / base
	}
	return Reverse(ret.String())
}
