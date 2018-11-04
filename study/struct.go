package study

import "fmt"

type Book struct {
	Title   string
	Author  string
	Subject string
	Id      int
}

func PrintBook(book Book) {
	fmt.Printf("Book title : %s\n", book.Title)
	fmt.Printf("Book author : %s\n", book.Author)
	fmt.Printf("Book subject : %s\n", book.Subject)
	fmt.Printf("Book book_id : %d\n", book.Id)
}
