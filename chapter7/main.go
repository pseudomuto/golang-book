package main

import (
	"fmt"
	"math"
)

func main() {
	num, even := halfEven(3)
	fmt.Println("Half of 3 is:", num, "and 3 = even number?:", even)

	num, even = halfEven(10)
	fmt.Println("Half of 10 is:", num, "and 10 = even number?:", even)

	fmt.Println("Largest in [1, 2, 4, 3]:", largestNumber(1, 2, 4, 3))

	oddNumber := makeOddGenerator()
	fmt.Println("Odd number:", oddNumber())
	fmt.Println("Odd number:", oddNumber())
	fmt.Println("Odd number:", oddNumber())
	fmt.Println("Odd number:", oddNumber())

	fmt.Println("Fib 6:", shittyFib(6))
}

func halfEven(x float64) (float64, bool) {
	return x / 2, math.Mod(x, 2) == 0
}

func largestNumber(values ...float64) float64 {
	largest := values[0]
	for _, value := range values {
		if value > largest {
			largest = value
		}
	}

	return largest
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
