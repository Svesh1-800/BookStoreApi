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
	}
	booksByte, _ := json.Marshal(books)
	w.Write(booksByte)
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
