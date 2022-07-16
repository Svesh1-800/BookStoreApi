package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/svesh1-800/bookstore/pkg/models"
	"github.com/svesh1-800/bookstore/pkg/utils"
)

func HandleGetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetBooks()
	if err != nil {
		log.Printf("Server Error%v\n", err)
		w.WriteHeader(500)
		w.Write(utils.JsonMessageByte("Internal server error"))
	} else {
		booksByte, _ := json.MarshalIndent(books, "", "\t")
		w.Write(booksByte)
	}
}

func HandleGetBookById(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	// get book id from URL
	bookId := query.Get("id")
	book, _, err := models.GetBookById(bookId)
	// send server error as response
	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(utils.JsonMessageByte("Internal server error"))
	} else {
		// check requested book exists or not
		if (models.Book{}) == book {
			w.Write(utils.JsonMessageByte("Book Not found"))
		} else {
			bookByte, _ := json.Marshal(book)
			w.Write(bookByte)
		}
	}
}

func HandleDeleteBookById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	bookId := query.Get("id")
	book, idx, err := models.GetBookById(bookId)
	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(utils.JsonMessageByte("Internal server error"))
	} else {
		if (models.Book{}) == book {
			w.Write(utils.JsonMessageByte("Book with the id doesn't exsist"))
		} else {
			books, err := models.GetBooks()
			if err != nil {
				log.Printf("Server Error%v\n", err)
				w.WriteHeader(500)
				w.Write(utils.JsonMessageByte("Internal server error"))

			} else {
				books = append(books[:idx], books[idx+1:]...)
				models.SaveBooks(books)
				w.Write(utils.JsonMessageByte("Book deleted successfully"))
			}
		}
	}
}
func HandleAddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		newBookByte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Client Error %v\n", err)
			w.WriteHeader(400)
			w.Write(utils.JsonMessageByte("Bad Request"))
		} else {
			books, _ := models.GetBooks()
			var newBooks models.Book
			json.Unmarshal(newBookByte, &newBooks)
			books = append(books, newBooks)

			err = models.SaveBooks(books)
			if err != nil {
				log.Printf("Server Error %v\n", err)
				w.WriteHeader(500)
				w.Write(utils.JsonMessageByte("Internal server error"))
			} else {
				w.Write(utils.JsonMessageByte("New book added successfully"))
			}
		}
	}
}
func HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		updateBookByte, err := ioutil.ReadAll(r.Body)
		// check for valid data from client
		if err != nil {
			log.Printf("Client Error %v\n", err)
			w.WriteHeader(400)
			w.Write(utils.JsonMessageByte("Bad Request"))
		} else {
			var updateBook models.Book // to update a book

			err = json.Unmarshal(updateBookByte, &updateBook) // new book added
			utils.CheckError(err)
			id := updateBook.Id

			book, _, _ := models.GetBookById(id)
			// check requested book exists or not
			if (models.Book{}) == book {
				w.Write(utils.JsonMessageByte("Book Not found"))
			} else {
				books, _ := models.GetBooks()

				for i, book := range books {
					if book.Id == updateBook.Id {
						books[i] = updateBook
					}
				}
				// write books in books.json
				err = models.SaveBooks(books)
				// send server error as response
				if err != nil {
					log.Printf("Server Error %v\n", err)
					w.WriteHeader(500)
					w.Write(utils.JsonMessageByte("Internal server error"))
				} else {
					w.Write(utils.JsonMessageByte("Book updated successfully"))
				}
			}
		}
	}
}
