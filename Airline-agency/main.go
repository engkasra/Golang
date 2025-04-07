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
		line := scanner.Text()
		parts := strings.Fields(line)

		country := parts[0]
		code := parts[1]

		CodeToCountry[code] = country

	}

	scanner.Scan()
	TelephonNumberCount, _ := strconv.Atoi(scanner.Text())
	var CodeCountries []string
	// var ContryName []string
	for i := 0; i < TelephonNumberCount; i++ {

		scanner.Scan()
		TelephonNumber := scanner.Text()
		code := TelephonNumber[:4]
		CodeCountries = append(CodeCountries, code)

		if country, ok := CodeToCountry[code]; ok {
			fmt.Println(country)

		}
	}

}
