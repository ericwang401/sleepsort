package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("magically sorting numbers")

	unsorted := [6]int{1, 3, 2, 5, 6, 4}
	sorted := []int{}

	for _, x := range unsorted {
		go func(x int, sorted *[]int) {
			time.Sleep(time.Duration(x) * time.Second)

			*sorted = append(*sorted, x)

			if len(*sorted) == 6 {
				fmt.Println(*sorted)
			}

		}(x, &sorted)
	}

	fmt.Scanln()
}
