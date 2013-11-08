package main

import "fmt"

func main() {
	oddNumber := makeOddGenerator()
	fmt.Println("Odd number:", oddNumber())
	fmt.Println("Odd number:", oddNumber())
	fmt.Println("Odd number:", oddNumber())
	fmt.Println("Odd number:", oddNumber())

	fmt.Println("Fib 6:", shittyFib(6))
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func shittyFib(n int) int {
	if n <= 1 {
		return n
	}

	return shittyFib(n - 1) + shittyFib(n - 2)
}
