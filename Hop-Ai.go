package main

import (
	"fmt"
)

func notmain() {
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			count := i / 3
			for j := 0; j < count; j++ {
				fmt.Print("Hope ")
			}
			fmt.Println() // move to next line after printing all "Hope"
		} else {
			fmt.Println(i)
		}
	}
}
