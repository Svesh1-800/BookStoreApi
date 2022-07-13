package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PORT string = ":8000"

type Message struct {
	Msg string
}

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

func jsonMessageByte(msg string) []byte {
	errMessage := Message{msg}
	byteContent, _ := json.Marshal(errMessage)
	return byteContent
}

func main() {
	http.HandleFunc("/books", handleGetBooks)

	http.HandleFunc("/book", handleGetBookById)

	http.HandleFunc("/delete", handleDeleteBookById)

	http.HandleFunc("/add", handleAddBook)

	http.HandleFunc("/update", handleUpdateBook)
	fmt.Printf("App is listening on %v\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	// stop the app is any error to start the server
	if err != nil {
		log.Fatal(err)
	}
}
func handleGetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := getBooks()
	if err != nil {
		log.Printf("Server Error%v\n", err)
		w.WriteHeader(500)
		w.Write(jsonMessageByte("Internal server error"))
	} else {
		booksByte, _ := json.MarshalIndent(books, "", "\t")
		w.Write(booksByte)
	}
}

func handleGetBookById(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	// get book id from URL
	bookId := query.Get("id")
	book, _, err := getBookById(bookId)
	// send server error as response
	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(jsonMessageByte("Internal server error"))
	} else {
		// check requested book exists or not
		if (Book{}) == book {
			w.Write(jsonMessageByte("Book Not found"))
		} else {
			bookByte, _ := json.Marshal(book)
			w.Write(bookByte)
		}
	}
}

func handleDeleteBookById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	bookId := query.Get("id")
	book, idx, err := getBookById(bookId)
	if err != nil {
		log.Printf("Server Error %v\n", err)
		w.WriteHeader(500)
		w.Write(jsonMessageByte("Internal server error"))
	} else {
		if (Book{}) == book {
			w.Write(jsonMessageByte("Book with the id doesn't exsist"))
		} else {
			books, err := getBooks()
			if err != nil {
				log.Printf("Server Error%v\n", err)
				w.WriteHeader(500)
				w.Write(jsonMessageByte("Internal server error"))

			} else {
				books = append(books[:idx], books[idx+1:]...)
				saveBooks(books)
				w.Write(jsonMessageByte("Book deleted successfully"))
			}
		}
	}
}
func handleAddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		newBookByte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Client Error %v\n", err)
			w.WriteHeader(400)
			w.Write(jsonMessageByte("Bad Request"))
		} else {
			books, _ := getBooks()
			var newBooks Book
			json.Unmarshal(newBookByte, &newBooks)
			books = append(books, newBooks)

			err = saveBooks(books)
			if err != nil {
				log.Printf("Server Error %v\n", err)
				w.WriteHeader(500)
				w.Write(jsonMessageByte("Internal server error"))
			} else {
				w.Write(jsonMessageByte("New book added successfully"))
			}
		}
	}
}
func handleUpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		updateBookByte, err := ioutil.ReadAll(r.Body)
		// check for valid data from client
		if err != nil {
			log.Printf("Client Error %v\n", err)
			w.WriteHeader(400)
			w.Write(jsonMessageByte("Bad Request"))
		} else {
			var updateBook Book // to update a book

			err = json.Unmarshal(updateBookByte, &updateBook) // new book added
			checkError(err)
			id := updateBook.Id

			book, _, _ := getBookById(id)
			// check requested book exists or not
			if (Book{}) == book {
				w.Write(jsonMessageByte("Book Not found"))
			} else {
				books, _ := getBooks()

				for i, book := range books {
					if book.Id == updateBook.Id {
						books[i] = updateBook
					}
				}
				// write books in books.json
				err = saveBooks(books)
				// send server error as response
				if err != nil {
					log.Printf("Server Error %v\n", err)
					w.WriteHeader(500)
					w.Write(jsonMessageByte("Internal server error"))
				} else {
					w.Write(jsonMessageByte("Book updated successfully"))
				}
			}
		}
	}
}
func checkError(err error) {
	if err != nil {
		log.Printf("Error - %v", err)
	}

}
