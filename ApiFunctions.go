package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"strconv"
)

func optionsRequest(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusOK)
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

type LoginAndRegister struct {
	Username 	string 		`json:"username"`
	Password 	string 		`json:"password"`
}

type remember_token struct {
	RemeberToken string 	`json:"remember_token"`
}

type User struct {
	Id 				int		`json:"id"`
	RoleId			int 	`json:"role_id"`
	Username 		string	`json:"username"`
	Password 		string	`json:"password"`
	RememberToken	string	`json:"remember_token"`
}

type Comment struct {
	Id 				int		`json:"id"`
	UserId			int		`json:"user_id"`
	Username		string	`json:"username"`
	BookId			int		`json:"book_id"`
	Text			string	`json:"comment_text"`
}

type Book struct {
	Id				int 	`json:"id"`
	Title			string	`json:"title"`
	Author			string	`json:"author"`
	Description		string	`json:"description"`
	Quantity		int		`json:"quantity"`
	RatingSum		int		`json:"rating_sum"`
	NumberRates		int		`json:"number_rates"`
	Rating			float32 `json:"rating"`
	ImageName		string	`json:"image_name"`
	PageNumber		int 	`json:"page_number"`
	Comment[]		Comment `json:"comments"`
}

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
}

type BookRating struct {
	Rate			int		`json:"rate"`
	BookID			int		`json:"bookID"`
}

type BookTaking struct {
	RememberToken	string	`json:"remember_token"`
	BookID			int		`json:"bookID"`
}

type GetComment struct {
	RememberToken	string	`json:"remember_token"`
	BookID			int		`json:"bookID"`
	Comment			string	`json:"comment"`
}

type ReservedBook struct {
	Id 				int		`json:"id"`
	Username		string	`json:"username"`
	Title			string	`json:"title"`
	ExpirationDate	string	`json:"expiration_date"`
}

func loginFunc(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	var login LoginAndRegister
	_ = json.NewDecoder(r.Body).Decode(&login)

	status, rememberToken := loginUser(login.Username, login.Password)

	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rememberToken)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func registerFunc(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	var register LoginAndRegister
	_ = json.NewDecoder(r.Body).Decode(&register)

	status := registerUser(register.Username, register.Password)

	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func getUserInfo(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	var rememberToken remember_token

	_ = json.NewDecoder(r.Body).Decode(&rememberToken)

	user,err := getUser(rememberToken.RemeberToken)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(user)
}

func addBook(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w,r)

	file, handler, err := r.FormFile("uploadfile")

	checkErr(err)
	defer file.Close()

	pageNumber, _ := strconv.Atoi(r.Form["page_number"][0])
	quantity, _   := strconv.Atoi(r.Form["quantity"][0])

	f, err := os.OpenFile("./images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	checkErr(err)

	defer f.Close()

	io.Copy(f, file)

	status := insertBook(r.Form["title"][0],r.Form["author"][0],r.Form["description"][0], quantity, handler.Filename, pageNumber)

	if status {
		http.Redirect(w,r, "http://localhost:8080/add-book", http.StatusSeeOther)
	} else {
		http.Redirect(w,r, "http://localhost:8080/add-book", http.StatusSeeOther)
	}

}

func getBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	setupResponse(&w,r)

	books, error := getAllBooks()

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error while querying")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}
}

func getSingleBook(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w,r)

	vars := mux.Vars(r)
	key := vars["id"]
	id,err := strconv.Atoi(key)
	checkErr(err)

	book, error := getBook(id)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
	}else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

func bookRating(w http.ResponseWriter, r *http.Request)  {

	setupResponse(&w,r)

	var bookRating BookRating
	_ = json.NewDecoder(r.Body).Decode(&bookRating)

	status := rateBook(bookRating.Rate, bookRating.BookID)

	if status {
		w.WriteHeader(http.StatusOK)
	}else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func takeBook(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w,r)

	var bookTaking BookTaking

	_ = json.NewDecoder(r.Body).Decode(&bookTaking)

	status := reserveBook(bookTaking.RememberToken, bookTaking.BookID)

	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getComment(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w,r)

	var comment GetComment

	_ = json.NewDecoder(r.Body).Decode(&comment)

	status := insertComment(comment.RememberToken, comment.BookID, comment.Comment)

	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getReservedBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	setupResponse(&w,r)

	reservedBooks,error := queryReservedBooks()

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
	}else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(reservedBooks)
	}
}

func setAsReturn(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w,r)

	vars := mux.Vars(r)
	key := vars["id"]
	id,err := strconv.Atoi(key)
	checkErr(err)

	status := setAsTrue(id)

	if status {
		w.WriteHeader(http.StatusOK)
	}else {
		w.WriteHeader(http.StatusBadRequest)
	}

}