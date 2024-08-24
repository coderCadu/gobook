package main

import (
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// & passa o endereço do ponteiro
// * recupera o valor do ponteiro ou define que aquela variavel será um ponteiro
func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)
	BookHandlers := web.NewBookHandlers(bookService)

	router := http.NewServeMux()
	router.HandleFunc("GET /books", BookHandlers.GetBooks)
	// router.HandleFunc("POST /books", BookHandlers.CreateBook)
	// router.HandleFunc("GET /books/{id}", BookHandlers.GetBookByID)
	// router.HandleFunc("PUT /books/{id}", BookHandlers.UpdateBook)
	// router.HandleFunc("DELETE /books/{id}", BookHandlers.DeleteBook)

	log.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", router)
}


// CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, author TEXT NOT NULL, genre TEXT NOT NULL);