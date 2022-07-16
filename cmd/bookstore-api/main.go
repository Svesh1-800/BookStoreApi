package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/svesh1-800/bookstore/pkg/routes"
)

const PORT string = ":8000"

func main() {
	f, _ := os.Getwd()
	fmt.Println(filepath.Base(f))
	fmt.Println(filepath.Dir(f))
	routes.RegisterBookStoreRoutes()
	fmt.Printf("App is listening on %v\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	// stop the app is any error to start the server
	if err != nil {
		log.Fatal(err)
	}
}
