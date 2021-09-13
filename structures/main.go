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
	fmt.Printf("Book 2 title : %s\n", book2.Title)
	fmt.Printf("Book 2 author : %s\n", book2.Author)
	fmt.Printf("Book 2 subject : %s\n", book2.Subject)
	fmt.Printf("Book 2 book_id : %d\n", book2.BookID)
}
