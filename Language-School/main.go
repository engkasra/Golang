package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Evaluation struct {
	Name   string
	Result string
}

func main() {
	//Part 1
	fmt.Print("How many Teacher there: ")
	scanner := bufio.NewScanner(os.Stdin)
	// what is doing scanner.scan and difference between bufio.newscanner?
	scanner.Scan()
	TeacherCount, _ := strconv.Atoi(scanner.Text())
	fmt.Println("%%%%% Your number %%%%%", TeacherCount)

	var results []Evaluation

	// Part 2

	for i := 0; i < TeacherCount; i++ {

		// Part 2.1

		fmt.Println("Name of the Teacher: ")
		// is requirement to using the os.stdin for string
		scanner.Scan()
		// what does doing scanner.text
		TeacherName := scanner.Text()
		fmt.Println("Your teacher's name is: ", TeacherName)

		// Part 2.2

		scanner.Scan()
		scorestr := strings.Fields(scanner.Text())

		// Part 2.3

		var sum int
		for _, s := range scorestr {

			score, _ := strconv.Atoi(s)
			sum += score
		}

		// Part 2.4

		count := len(scorestr)
		avg := sum / count
		println("Your Average:", avg)

		// Part 2.5

		var result string

		if avg >= 80 {
			result = "Excellent"
		} else if avg >= 60 {
			result = "Very Good"
		} else if avg >= 40 {
			result = "Good"
		} else {
			result = "Fair"
		}

		results = append(results, Evaluation{Name: TeacherName, Result: result})
		fmt.Printf("Here is state of result in array: %+v", results)

	}
	fmt.Println("\nEvaluation Results: ")
	for _, r := range results {

		fmt.Printf("%s %s\n", r.Name, r.Result)
	}

}
