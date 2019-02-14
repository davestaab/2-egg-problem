package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[2])
	v, _ := strconv.Atoi(os.Args[1])
	g := int(math.Ceil(math.Sqrt(float64(v))))
	i := 0

	f := func(p int) {
		for i < n {
			i += p
			fmt.Printf("%d\n%t\n", i, i >= n)
		}
	}
	f(g)
	i -= g
	f(1)
	fmt.Printf("Floor %d", i)
}
