package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("How many Telephone Number there is? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	CountryCount, _ := strconv.Atoi(scanner.Text())
	var CodeToCountry = make(map[string]string)
	for i := 0; i < CountryCount; i++ {

		scanner.Scan()
		// line := scanner.Text()
		// parts := strings.Fields(line)

		parts := strings.Fields(scanner.Text())
		country := parts[0]
		code := parts[1]

		CodeToCountry[code] = country

	}

	var results []string
	scanner.Scan()
	TelephonNumberCount, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < TelephonNumberCount; i++ {

		scanner.Scan()
		number := scanner.Text()

		found := false
		for code, country := range CodeToCountry {

			if strings.HasPrefix(number, code) {
				results = append(results, country)
				found = true
				break
			}

		}
		if !found {
			results = append(results, "Invalid Number")
		}

		// Following comment is working but very simple!
		// if len(number) >= 4 {
		// 	code := number[:4]

		// 	if country, ok := CodeToCountry[code]; ok {
		// 		fmt.Println(country)

		// 	} else {
		// 		fmt.Println("Invalid Number")
		// 	}
		// } else {
		// 	fmt.Println("Invalid Number")
		// }

	}
	for _, res := range results {
		fmt.Println(res)
	}

}
