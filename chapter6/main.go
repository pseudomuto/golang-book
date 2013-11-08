package main

import "fmt"

func main() {
	// average()
	smallest()
}

func average() {
	// var x [5]float64
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83
	x := []float64 { 98, 93, 77, 82, 83 }

    var total float64 = 0
    // for i := 0; i < len(x); i++ {
    // 	total += x[i]
    // }
    for _, value := range x {
    	total += value
    }

    fmt.Println("Average x:", total / float64(len(x)))
}

func smallest() {
	x := []int {
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	smallestValue := x[0]
	for _, value := range x {
		if value < smallestValue {
			smallestValue = value
		}
	}

	fmt.Println("Smallest x:", smallestValue)
}
