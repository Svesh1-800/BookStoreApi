package routes

import (
	"log"
	"net/http"

	"github.com/svesh1-800/bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func() {

	log.Println("we are here")

	http.HandleFunc("/books", controllers.HandleGetBooks)

	http.HandleFunc("/book", controllers.HandleGetBookById)

	http.HandleFunc("/delete", controllers.HandleDeleteBookById)

	http.HandleFunc("/add", controllers.HandleAddBook)

	http.HandleFunc("/update", controllers.HandleUpdateBook)

	log.Println("we end")
}
