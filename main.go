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
func getBooks() ([]Book, error) {
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
func getBookById(id string) (Book, int, error) {
	books, err := getBooks()
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
func checkError(err error) {
	if err != nil {
		log.Printf("Error - %v", err)
	}

}
func saveBooks(books []Book) error {

	// converting into bytes for writing into a file
	booksBytes, err := json.Marshal(books)

	checkError(err)

	err = ioutil.WriteFile("./books.json", booksBytes, 0644)

	return err

}
