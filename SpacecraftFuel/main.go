package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countArithmeticSequences(nums []int) int {

	count := 0
	n := len(nums)

	for i := 0; i < n-2; i++ {

		d := nums[i+1] - nums[i]
		lenght := 2

		for j := i + 2; j < n; j++ {

			if nums[j]-nums[j-1] == d {
				lenght++
				if lenght >= 3 {

					count++

				}
			} else {
				break
			}
		}

	}
	return count
}

type Result struct {
	Name  string
	Count int
}

func main() {

	fmt.Println("How many Spaccraft there?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Spacecraftcount, _ := strconv.Atoi(scanner.Text())
	fmt.Println("for just using", Spacecraftcount)

	var results []Result

	for i := 0; i < Spacecraftcount; i++ {

		fmt.Println("Now import the name and fuel:")
		scanner.Scan()
		Spacecraft_fullpart := scanner.Text()
		Spacecraft_fields := strings.Fields(Spacecraft_fullpart)
		Spacecraft_name := Spacecraft_fields[0]
		fmt.Println("Here is Part 0 :", Spacecraft_name)

		fuelStrs := Spacecraft_fields[1:]
		fmt.Printf("Here is the fuel strings %v\n", fuelStrs)
		var fuelInts []int
		for _, s := range fuelStrs {

			num, _ := strconv.Atoi(s)
			fuelInts = append(fuelInts, num)
		}

		fmt.Println("Here is the Fuelnumbers:", fuelInts)
		fuelCount := countArithmeticSequences(fuelInts)

		results = append(results, Result{Name: Spacecraft_name, Count: fuelCount})
		fmt.Println("THIS IS results Array: ", results, "Arithmetic count:", fuelCount)
	}
	sort.Slice(results, func(i, j int) bool {

		if results[i].Count == results[j].Count {
			return results[i].Name < results[j].Name
		}
		return results[i].Count > results[j].Count

	})
	for _, r := range results {

		fmt.Printf("Final Result: %s %d\n", r.Name, r.Count)
	}
}
