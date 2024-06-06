package main

import "fmt"

func Atoi(str string) int {
	var result int
	sign := 1
	firstNoneSpace := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}
		if str[i] == '-' {
			sign = -1
			firstNoneSpace = i
			break
		}
	}
	for i := firstNoneSpace; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		result = result*10 + int(str[i]-'0')
	}
	return result * sign
}

func main() {
	n := "400"
	fmt.Println(Atoi(n))
}
