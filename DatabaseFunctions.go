package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func HashString(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func openDB()  *sql.DB{
	db, err := sql.Open("mysql","dario:notImportant@/e-library")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func checkErrorWithReturn(err error) bool{
	if err != nil {
		panic(err.Error())
		return false
	}

	return true
}

func loginUser(username string, password string) (bool, string) {
	db := openDB()
	var user User

	err := db.QueryRow("SELECT id,role_id,username,password,remember_token FROM users WHERE username=?",username).Scan(&user.Id, &user.RoleId ,&user.Username, &user.Password, &user.RememberToken)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	}

	match :=CheckPasswordHash(password, user.Password)

	return match, user.RememberToken
}

func registerUser(username string, password string) bool {
	db := openDB()
	stmIns, err := db.Prepare("INSERT INTO users(username, password, remember_token) VALUES(?,?,?)")
	checkErr(err)

	rememberToken, _ := HashString(RandStringRunes(20))

	defer stmIns.Close()
	hashPassword, _ := HashString(password)

	_ , err = stmIns.Exec(username, hashPassword, rememberToken)

	return checkErrorWithReturn(err)
}

func getUser(rememberToken string) (User, error) {
	db := openDB()
	var user User

	err := db.QueryRow("SELECT id,role_id,username,password,remember_token FROM users WHERE remember_token=?", rememberToken).Scan(&user.Id, &user.RoleId ,&user.Username, &user.Password, &user.RememberToken)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
	}

	return user,err
}

func insertBook(title string, author string, description string, quantity int, imageName string ,pageNumber int) (bool) {

	db := openDB()

	stmIns, err := db.Prepare("INSERT INTO books(title,author, description, quantity, image_name,page_number) VALUES(?,?,?,?,?,?)")
	checkErr(err)

	defer stmIns.Close()

	_ , err = stmIns.Exec(title,author,description,quantity,imageName,pageNumber)

	return checkErrorWithReturn(err)
}

func getAllBooks() ([]Book, error) {

	db := openDB()

	var book Book
	var books []Book

	rows,err := db.Query("SELECT id,title,author,description,quantity,rating_sum,number_rates,rating,image_name,page_number FROM books")
	checkErr(err)

	for rows.Next() {
		rows.Scan(&book.Id,&book.Title,&book.Author,&book.Description,&book.Quantity,&book.RatingSum,&book.NumberRates,&book.Rating,&book.ImageName,&book.PageNumber)

		books = append(books,book)
	}

	defer db.Close()

	return books, err
}

func getUsername(id int) string {

	db := openDB()

	var user User

	err := db.QueryRow("SELECT id,role_id,username,password,remember_token FROM users WHERE id=?",id).Scan(&user.Id, &user.RoleId ,&user.Username, &user.Password, &user.RememberToken)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
	}

	return user.Username
}

func getBook(id int) (Book, error) {

	db := openDB()

	var book Book
	var comment Comment
	var comments []Comment


	err := db.QueryRow("SELECT id,title,author,description,quantity,rating_sum,number_rates,rating,image_name,page_number FROM books WHERE id=?", id).Scan(&book.Id, &book.Title ,&book.Author, &book.Description, &book.Quantity, &book.RatingSum, &book.NumberRates, &book.Rating ,&book.ImageName, &book.PageNumber)
	rows, commErr := db.Query("SELECT id,user_id,book_id,text FROM comments WHERE book_id=?",id)
	checkErr(commErr)

	for rows.Next() {
		rows.Scan(&comment.Id,&comment.UserId,&comment.BookId,&comment.Text)
		comment.Username = getUsername(comment.UserId)
		comments = append(comments,comment)
	}

	book.Comment = comments

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
	}

	return book,err
}

func rateBook(rate int, bookID int) bool {

	db := openDB()

	stmt, err := db.Prepare("UPDATE books SET rating_sum=(SELECT rating_sum WHERE id=?)+?, number_rates=(SELECT number_rates WHERE id=?)+1 WHERE id=?")

	result,err := stmt.Exec(bookID,rate,bookID, bookID)

	ratingStmt, err := db.Prepare("UPDATE books SET rating=(SELECT rating_sum WHERE id=?)/(SELECT number_rates WHERE id=?) WHERE id=?")

	ratingResult,err := ratingStmt.Exec(bookID,bookID,bookID)


	defer stmt.Close()

	defer ratingStmt.Close()

	var numResults, _ = result.RowsAffected()
	var ratingResults, _ = ratingResult.RowsAffected()

	if err != nil || numResults != 1 && ratingResults != 1 {
		return false
	}

	return true
}

func decrementQuantity(bookID int) bool {

	db := openDB()

	stmt, err := db.Prepare("UPDATE books SET quantity=(SELECT quantity WHERE id=?)-1 WHERE id=?")

	result,err := stmt.Exec(bookID,bookID)

	defer stmt.Close()

	var numResults, _ = result.RowsAffected()

	if err != nil || numResults != 1 {
		return false
	}

	return true
}

func reserveBook(remember_token string, bookID int) bool {

	db := openDB()

	t := time.Now()
	t2 := t.AddDate(0, 1, 0)

	user, _ := getUser(remember_token)

	stmIns, err := db.Prepare("INSERT INTO book_taking(user_id,book_id,expiration_date) VALUES(?,?,?)")
	checkErr(err)

	defer stmIns.Close()

	_, err = stmIns.Exec(user.Id,bookID,t2)

	return decrementQuantity(bookID) && checkErrorWithReturn(err)

}

func insertComment(remember_token string,bookID int, comment string) bool {

	db := openDB()

	user, _ := getUser(remember_token)

	stmIns, err := db.Prepare("INSERT INTO comments(user_id,book_id,text) VALUES(?,?,?)")
	checkErr(err)

	defer stmIns.Close()

	_, err = stmIns.Exec(user.Id,bookID,comment)

	return checkErrorWithReturn(err)
}

func queryReservedBooks() ([]ReservedBook, error){

	db := openDB()

	var reservedBook ReservedBook
	var reservedBooks []ReservedBook

	query := `
		SELECT bt.id,username,title,expiration_date
		FROM book_taking bt 
		INNER JOIN users u ON bt.user_id = u.id
		INNER JOIN books b ON bt.book_id = b.id
		WHERE returned = 0`

	rows, err := db.Query(query)
	checkErr(err)

	defer db.Close()

	for rows.Next() {
		rows.Scan(&reservedBook.Id,&reservedBook.Username, &reservedBook.Title, &reservedBook.ExpirationDate)

		reservedBooks = append(reservedBooks, reservedBook)
	}

	return reservedBooks,err
}

func setAsTrue(id int) bool {

	db := openDB()

	stmt, err := db.Prepare("UPDATE book_taking SET returned=1 WHERE id=?")

	result,err := stmt.Exec(id)

	defer stmt.Close()

	var numResults, _ = result.RowsAffected()

	if err != nil || numResults != 1  {
		return false
	}

	return true
}
