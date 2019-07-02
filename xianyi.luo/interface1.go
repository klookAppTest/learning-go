package main

import "fmt"

// 定义接口
type Describer interface {
	describe()
}

type Person struct {
	name string
	age int
}

func (p Person) describe(){
	fmt.Printf("%s is %d years old\n",p.name,p.age)
}

type Address struct {
	state string
	country string
}

// 该函数传入参数为指针类型，因此main函数中传入必须为地址类型
func (a *Address) describe(){
	fmt.Printf("State %s Country %s\n",a.state,a.country)
}


func main(){
	var d1 Describer
	p1 := Person{"Sam",24}
	d1 = p1
	d1.describe()

	p2 := Person{"James",31}
	d1 = &p2
	d1.describe()

	var d2 Describer
	a := Address{"Washington","USA"}
	d2 = &a
	d2.describe()
	a.describe()
}