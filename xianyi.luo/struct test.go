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

	fmt.Println(Book1)

	change_title(&Book1)

	fmt.Println(Book1)

}

// 结构体作为指针传入函数，改变结构体的参数数据
func change_title(book *Books){
	book.title = "chaged_title"
}