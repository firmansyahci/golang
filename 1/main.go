package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Book struct {
	Id       string
	Kode     string
	Category string
	Name     string
	Height   string
}

func main() {

	fmt.Println("input book : ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	s := strings.Split(line, " ")

	var book []Book
	var category string

	for _, value := range s {
		switch value[0] {
		case 48: // 0
			category = "General (000)"
		case 49: // 1
			category = "Philosophy, Psychology (100)"
		case 50: // 2
			category = "Religion (200)"
		case 51: // 3
			category = "Social Sciences (300)"
		case 52: // 4
			category = "Language (400)"
		case 53: // 5
			category = "Science, Math (500)"
		case 54: // 6
			category = "Applied Sciences (600)"
		case 55: // 7
			category = "Arts (700)"
		case 56: // 8
			category = "Literature (800)"
		case 57: // 9
			category = "Geography, History (900)"
		default:
			fmt.Println("Hi")
		}

		book = append(book, Book{Id: value, Kode: string(value[0]), Category: category, Name: string(value[1]), Height: string(value[2:4])})
	}

	sort.Slice(book, func(p, q int) bool {
		switch strings.Compare(book[p].Category, book[q].Category) {
		case -1:
			return true
		case 1:
			return false
		}
		return book[p].Height > book[q].Height
	})

	var currentKode, currentName string
	countSame := true

	for _, value := range book {
		if value.Kode != currentKode {
			fmt.Printf("%s ", value.Id)
		} else {
			if value.Name == currentName {
				if countSame {
					fmt.Printf("%s ", value.Id)
					countSame = false
				}
			} else {
				fmt.Printf("%s ", value.Id)
				countSame = true
			}
		}
		currentKode = value.Kode
		currentName = value.Name
	}
}
