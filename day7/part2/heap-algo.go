package main

import "fmt"

func generate(k int, numbers []int) {
	if k == 1 {
		fmt.Println(numbers)
		return
	}

	generate(k-1, numbers)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			numbers[i], numbers[k-1] = numbers[k-1], numbers[i]
		} else {
			numbers[0], numbers[k-1] = numbers[k-1], numbers[0]
		}
		generate(k-1, numbers)
	}
}
