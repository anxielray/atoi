package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	fmt.Println(Atoi(os.Args[1]))
}

func Atoi(arg string) (n int) {
	digit, number := 1, 0
	if strings.HasPrefix(arg, "-") {
		arg = string(arg[1:])
		for x := len(arg) - 1; x >= 0; x-- {
			number = int(rune(arg[x]) - 48)
			n += (digit * number)
			digit *= 10
		}
		return -n
	} else if strings.HasPrefix(arg, "+") {
		arg = arg[1:]
		for x := len(arg) - 1; x >= 0; x-- {
			number = int(rune(arg[x]) - 48)
			n += (digit * number)
			digit *= 10
		}
	} else {
		for x := len(arg) - 1; x >= 0; x-- {
			number = int(rune(arg[x]) - 48)
			n += (digit * number)
			digit *= 10
		}
	}

	return
}
