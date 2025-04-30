package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func AddBook(books map[int]string, isbn int, title string) {

	books[isbn] = title
}

func RemoveBook(books map[int]string, isbn int) {

	delete(books, isbn)
}

// var action string

// var Books = map[int]string{
// 	123: "Stevene Hawking",
// 	456: "Bill Gates",
// }

type BookEntry struct {
	ISBN  int
	Title string
}

var entries []BookEntry

func main() {

	var books = make(map[int]string)

	fmt.Println("This is begin!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	CountAction, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < CountAction; i++ {

		scanner.Scan()
		parts := strings.Fields(scanner.Text())

		if parts[0] == "ADD" {

			isbn, _ := strconv.Atoi(parts[1])
			title := strings.Join(parts[2:], " ")
			AddBook(books, isbn, title)

		} else if parts[0] == "REMOVE" {

			isbn, _ := strconv.Atoi(parts[1])
			RemoveBook(books, isbn)

		}
	}
	for isbn, title := range books {

		entries = append(entries, BookEntry{ISBN: isbn, Title: title})
	}
	sort.Slice(entries, func(i, j int) bool {

		if entries[i].Title == entries[j].Title {
			return entries[i].ISBN < entries[j].ISBN
		}
		return entries[i].Title < entries[j].Title

	})

	for _, entry := range entries {

		fmt.Println(entry.ISBN)
	}
}
