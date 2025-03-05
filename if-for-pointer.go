package main

import "fmt"

func incrementByTen(num *int) {
	*num += 10
}

func ExampleIfForPointer() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original numbers:", numbers)
	for i, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}

		incrementByTen(&numbers[i])
	}

	fmt.Println("Modified numbers:", numbers)
}
