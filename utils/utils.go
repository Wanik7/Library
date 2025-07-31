package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Book struct {
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Year   int      `json:"year"`
	Genres []string `json:"genres"`
}

func LoadBooks() ([]Book, error) { // i use this function like sub-function in other functions
	data, err := os.ReadFile("utils/library.json")
	if err != nil {
		return []Book{}, errors.New("read file")
	}
	var booksVault []Book
	err = json.Unmarshal(data, &booksVault)
	if err != nil {
		return []Book{}, errors.New("parse from json")
	}

	return booksVault, nil
}

func AddBook(title string, author string, year int, genres []string) error {
	switch {
	case title == "":
		return errors.New("book title cannot be empty")
	case author == "":
		return errors.New("book author cannot be empty")
	case year <= 0 || year > time.Now().Year():
		return errors.New("invalid year")
	case len(genres) == 0:
		return errors.New("book must contains at least one genre")
	default:
		newBook := Book{
			Title:  title,
			Author: author,
			Year:   year,
			Genres: genres,
		}

		books, err := LoadBooks()
		if err != nil {
			return errors.New(err.Error())
		}

		for _, v := range books { // check for duplicates
			if v.Author == newBook.Author && v.Title == newBook.Title && v.Year == newBook.Year {
				return errors.New("this book already exists in library")
			}
		}

		books = append(books, newBook)
		toJson, _ := json.MarshalIndent(books, "", "  ")
		os.WriteFile("utils/library.json", toJson, 0644)
		
		fmt.Println("The book was append")
		return nil
	}
}

func FindBooksByAuthor(author string) ([]Book, error) {
	books, err := LoadBooks()
		if err != nil {
			return []Book{}, errors.New(err.Error())
		}

	var booksByAuthor []Book
	for _, v := range books {
		if strings.Contains(v.Author, author) {
			booksByAuthor = append(booksByAuthor, v)
		}
	}
	if len(booksByAuthor) == 0 {
		return []Book{}, errors.New("there is no books")
	}
	return booksByAuthor, nil
}

func DeleteBookByTitle(title string) error {
	books, err := LoadBooks()
	if err != nil {
		return errors.New(err.Error())
	}

	var newBooks []Book
	for _, v := range books {
		if v.Title != title {
			newBooks = append(newBooks, v)
		}
	}
	if len(newBooks) == len(books) {
		return errors.New("this book is not in library")
	}

	toJson, _ := json.MarshalIndent(newBooks, "", "  ")
	os.WriteFile("utils/library.json", toJson, 0644)

	fmt.Println("The book was removed")
	return nil
}

func FindBookByGenre(genre string) ([]Book, error) {
	books, err := LoadBooks()
	if err != nil {
		return []Book{}, errors.New(err.Error())
	}

	var newBooks []Book
	for _, book := range books {
		for _, genres := range book.Genres {
			if genres == genre {
				newBooks = append(newBooks, book)
			}
		}
	}
	if len(newBooks) == 0 {
		return []Book{}, errors.New("there is no book with this genre")
	}

	return newBooks, nil
}