package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
)

type BookHandlers struct {
	bookService *service.BookService
}

func NewBookHandlers(bookService *service.BookService) *BookHandlers {
	return &BookHandlers{bookService: bookService}
}

func (handler *BookHandlers) GetBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := handler.bookService.GetBooks()
	if err != nil {
		http.Error(writer, "Failed to get books", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}