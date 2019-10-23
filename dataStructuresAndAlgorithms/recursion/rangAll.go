package recursion

import "fmt"

// Ran prints a slice's all rang
func Ran(slice []int) {
	_ran(slice, 0)
}

func _ran(slice []int, start int) {
	len := len(slice)
	if start == len-1 {
		fmt.Printf("%v \n", slice)
	}
	for i := start; i < len; i++ {
		if i == start {
			_ran(slice, start+1)
		} else if slice[i] != slice[start] {
			slice[start], slice[i] = slice[i], slice[start]
			_ran(slice, start+1)
			slice[start], slice[i] = slice[i], slice[start]
		}
	}
}
