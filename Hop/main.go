package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter multiple and max number (e.g. 3 10): ")
	scanner.Scan()
	line := scanner.Text()

	parts := strings.Fields(line)
	if len(parts) != 2 {
		fmt.Print("Please enter exactly two numbers.")
		return
	}

	multiple, err1 := strconv.Atoi(parts[0])
	max, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. Please enter numbers.")
		return
	}

	for i := 1; i <= max; i++ {

		if i%multiple == 0 {
			count := i / multiple
			for j := 0; j < count; j++ {
				fmt.Print("Hope ")
			}
			fmt.Println()

		} else {
			fmt.Println(i)
		}

	}

}
