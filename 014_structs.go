package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	/* print Book1 info */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* print Book2 info */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	/* print Book info from a function */
	printBook(Book1)
	printBook(Book2)

	/* print book info from a function with pointer */
	printBookByPointer(&Book1)
	printBookByPointer(&Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title from fn : %s\n", book.title)
	fmt.Printf("Book author from fn : %s\n", book.author)
	fmt.Printf("Book subject from fn : %s\n", book.subject)
	fmt.Printf("Book book_id from fn : %d\n", book.book_id)
}

func printBookByPointer(book *Books) {
	book.author = book.author + " & Me"

	fmt.Printf("Book title from ptr fn : %s\n", book.title)
	fmt.Printf("Book author from ptr fn : %s\n", book.author)
	fmt.Printf("Book subject from ptr fn : %s\n", book.subject)
	fmt.Printf("Book book_id from ptr fn : %d\n", book.book_id)
}
