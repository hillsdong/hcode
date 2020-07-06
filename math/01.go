package math

import "fmt"

//Q01 求一个整数的二进制表示
func Q01(decimal int64) string {
	var ret string
	var b int64
	for i := 32; i > 0; i = i - 1 {
		b = decimal & 1
		ret = fmt.Sprint(b) + ret

		decimal = decimal >> 1

	}

	return ret
}
