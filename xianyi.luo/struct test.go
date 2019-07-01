package main

import "fmt"


type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main()  {
	fmt.Println(Books{"Go 语言","www.klook.com","Go 语言教程",6495407})
	var Book1 Books

	Book1.title = "Go 语言"
	Book1.author = "www.klook.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	fmt.Println(Book1.title, Book1.author , Book1.subject , Book1.book_id)
}