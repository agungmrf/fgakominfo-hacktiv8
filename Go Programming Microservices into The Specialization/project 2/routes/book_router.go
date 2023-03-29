package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"project-2/controller"
)

type BookRouter interface {
	SetupRoutes(subrouter *mux.Router)
}

type bookRouter struct {
	controller controller.BookController
}

func NewBookRouter(controller controller.BookController) BookRouter {
	return &bookRouter{
		controller: controller,
	}
}

func (router *bookRouter) SetupRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("", router.getAllBooks).Methods(http.MethodGet)
	subrouter.HandleFunc("/{id}", router.getBookById).Methods(http.MethodGet)
	subrouter.HandleFunc("", router.createBook).Methods(http.MethodPost)
	subrouter.HandleFunc("/{id}", router.updateBook).Methods(http.MethodPut)
	subrouter.HandleFunc("/{id}", router.deleteBook).Methods(http.MethodDelete)
}

func (router *bookRouter) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := router.controller.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (router *bookRouter) getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := router.controller.GetBookById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (router *bookRouter) createBook(w http.ResponseWriter, r *http.Request) {
	var bookInput controller.BookInput
	err := json.NewDecoder(r.Body).Decode(&bookInput)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	book, err := router.controller.CreateBook(&bookInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (router *bookRouter) updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	var bookInput controller.BookInput
	err = json.NewDecoder(r.Body).Decode(&bookInput)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	book, err := router.controller.UpdateBook(uint(id), &bookInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (router *bookRouter) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	err = router.controller.DeleteBook(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})
}
