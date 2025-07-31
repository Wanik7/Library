package main

import (
	"fmt"
	"library/utils"
	"log"
	"strings"
)

func main() {
	// if you want to add a book
	err := utils.AddBook("Alexander Pushkin", "", 1823, []string{"something"})
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	// if you want to see all books in the library
	books, err := utils.LoadBooks()
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	for _, book := range books {
		var genres string
		for _, genre := range book.Genres {
			genres += genre + ", "
		}
		genres = strings.TrimRight(genres, ", ")
		fmt.Printf("Title: %s; Author: %s; Release year: %d; Genres: %s\n", book.Title, book.Author, book.Year, genres)
	}

	// if you want to find certain author books
	tolkiensBooks, err := utils.FindBooksByAuthor("Tolkien")
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	for _, book := range tolkiensBooks {
		var genres string
		for _, genre := range book.Genres {
			genres += genre + ", "
		}
		genres = strings.TrimRight(genres, ", ")
		fmt.Printf("Title: %s; Author: %s; Release year: %d; Genres: %s\n", book.Title, book.Author, book.Year, genres)
	}

	// if you want to delete book by its title
	err = utils.DeleteBookByTitle("The Hobbit")
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}

	// if you want to find book by certain genre
	booksByGenre, err := utils.FindBookByGenre("fantasy")
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	for _, book := range booksByGenre {
		var genres string
		for _, genre := range book.Genres {
			genres += genre + ", "
		}
		genres = strings.TrimRight(genres, ", ")
		fmt.Printf("Title: %s; Author: %s; Release year: %d; Genres: %s\n", book.Title, book.Author, book.Year, genres)
	}
	
}
