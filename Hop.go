package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Scanf("%d", &number)
	for i := 1; i <= number; i++ {

		if i%3 == 0 {
			count := i / 3
			for j := 0; j < count; j++ {
				fmt.Print("Hope ")
			}
			fmt.Println()

		} else {
			fmt.Println(i)
		}

	}

}
