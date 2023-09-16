package main

import "fmt"

func print(n int) {
	if n == 1 {
		fmt.Println(1)
		return
	}
	print(n - 1)
	fmt.Println(n)
}

func main() {
	print(7)
}
