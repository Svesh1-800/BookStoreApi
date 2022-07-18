package routes

import (
	"net/http"

	"github.com/svesh1-800/bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func() {

	http.HandleFunc("/books", controllers.HandleGetBooks)

	http.HandleFunc("/book", controllers.HandleGetBookById)

	http.HandleFunc("/delete", controllers.HandleDeleteBookById)

	http.HandleFunc("/add", controllers.HandleAddBook)

	http.HandleFunc("/update", controllers.HandleUpdateBook)

}
