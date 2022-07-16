package models

import (
	"encoding/json"
	"io/ioutil"

	"github.com/svesh1-800/bookstore/pkg/utils"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

func SaveBooks(books []Book) error {

	booksBytes, err := json.Marshal(books)
	utils.CheckError(err)

	err = ioutil.WriteFile("./books.json", booksBytes, 0644)

	return err

}
func GetBooks() ([]Book, error) {
	books := []Book{}
	booksByte, err := ioutil.ReadFile("./books.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(booksByte, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func GetBookById(id string) (Book, int, error) {
	books, err := GetBooks()
	var requestedBook Book
	var requestedBookIndex int

	if err != nil {
		return Book{}, 0, err
	}

	for i, book := range books {
		if book.Id == id {
			requestedBook = book
			requestedBookIndex = i
		}
	}

	return requestedBook, requestedBookIndex, nil
}
