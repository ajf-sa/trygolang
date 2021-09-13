package main

import "fmt"

type Book struct {
	Title   string
	Author  string
	Subject string
	BookID  int
}

func main() {

	var book1 Book
	var book2 Book

	book1.Title = "Telecom Billing"
	book1.Author = "Zara ali"
	book1.Subject = "Telecom Billing Tut"
	book1.BookID = 23423
	book2.Title = "Go Programing"
	book2.Author = "Mahesh ali"
	book2.Subject = "Go Programing Tut"
	book2.BookID = 2345

	fmt.Printf("Book 1 title : %s\n", book1.Title)
	fmt.Printf("Book 1 author : %s\n", book1.Author)
	fmt.Printf("Book 1 subject : %s\n", book1.Subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.BookID)

	/* print Book2 info */
	pringBook(book2)
}

func pringBook(book Book) {
	fmt.Printf("Book  title : %s\n", book.Title)
	fmt.Printf("Book  author : %s\n", book.Author)
	fmt.Printf("Book  subject : %s\n", book.Subject)
	fmt.Printf("Book  book_id : %d\n", book.BookID)
}
