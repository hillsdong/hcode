package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
		"   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}

func main() {
	var buf bytes.Buffer
	var nums string
	var bar bool
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s [-b|--bar] <whole-number> \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if len(os.Args) == 3 && (os.Args[1] == "-b" || os.Args[1] == "--bar") {
		bar = true
		nums = os.Args[2]
	} else {
		nums = os.Args[1]
	}

	var barlen int

	for i := range bigDigits[0] {
		for _, num := range nums {
			j := num - '0'
			if int(j) < len(bigDigits) {
				buf.WriteString(bigDigits[j][i])
				buf.WriteString(" ")
				if i == 0 {
					barlen += len(bigDigits[j][i]) + 1
				}
			} else {
				log.Fatal("invalid whole number")
			}
		}
		if i != len(bigDigits[0])-1 {
			buf.WriteString("\n")
		}
	}

	if bar {
		barlen--
		fmt.Println(strings.Repeat("*", barlen))
	}
	fmt.Println(buf.String())
	if bar {
		fmt.Println(strings.Repeat("*", barlen))
	}
}
