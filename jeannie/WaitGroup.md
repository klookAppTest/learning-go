#### 组合

```
// website.go
package website

import (
	"blog/post"
	"fmt"
)

type Website struct {
	Posts []post.Post
}

func New(posts []post.Post) Website {
	w := Website{posts}
	return w
}

func (w Website) Contents() {
	fmt.Println("contents of Website\n")
	for _, v := range w.Posts {
		v.Details()
		fmt.Println()
	}
}
```

```
//post.go
package post

import (
	"blog/author"
	"fmt"
)

type Post struct {
	Title string
	Content string
	author.Author
}

func New(title string, content string, a author.Author) Post {
	p := Post{title, content, a}
	return p
}

func (p Post) Details() {
	fmt.Println("Title: ", p.Title)
	fmt.Println("content: ", p.Content)
	fmt.Println("fullName: ", p.Author.FullName())
	fmt.Println("Bio: ", p.Author.Bio)
}
```
```
//author.go
package author

import "fmt"

type Author struct {
	FirstName string
	LastName string
	Bio string
}


func New(firstname string, lastname string, bio string) Author {
	a := Author{firstname, lastname, bio}
	return a
}

func (a Author) FullName() string {
	return fmt.Sprintf("fullname is: %s %s", a.FirstName, a.LastName)
}
```

```
package main

import (
	"blog/author"
	"blog/post"
	"blog/website"
)

// 组合取代继承
// Go 不支持继承，但它支持组合（Composition）。组合一般定义为“合并在一起”。汽车就是一个关于组合的例子：
// 一辆汽车由车轮、引擎和其他各种部件组合在一起。

// 通过嵌套结构体进行组合

func main() {
	a := author.New("jimmy", "hu", "text")
	//b := author.New("joli", "ma", "story")
	b := author.Author{
		FirstName: "joli",
		LastName: "ma",
		Bio: "story",
	}
	//p1 := post.New("title", "diff", a)
	p1 := post.Post{
		Title: "title1",
		Content: "content2",
		Author: a,
	}
	p2 := post.New("title2", "contnt2", a)
	p3 := post.New("title3", "content", a)
	p4 := post.New("title", "content", b)


	//p.Details()
	//w := website.Website{
	//	Posts: []post.Post{p1, p2, p3, p4},
	//}

	w := website.New([]post.Post{p1, p2, p3, p4})
	w.Contents()
}

```