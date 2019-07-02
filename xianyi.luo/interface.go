package main

import "fmt"

type SalaryCalculate interface {
	calculateSalary() int
}

// 类似于Java中的一个类，定义成员变量，只不过没有方法
type Employee struct {
	empId int
	basicSal int
}

type seniorEmployee struct {
	empId int
	basicSal int
	prize int
}

// 类似于Java类中的方法，并且一定要调用上面接口中的方法，不然报错，类似override，传入参数为结构体
func (se seniorEmployee) calculateSalary() int{
	return se.basicSal + se.prize
}

func (e Employee) calculateSalary() int{
	return e.basicSal
}

// 传入一个数组类型的函数包含上述的两个结构体，用range来进行遍历都调用接口中的函数
func totalExpanse(s []SalaryCalculate){
	expense := 0
	for _,v := range s{
		expense += v.calculateSalary()
	}
	fmt.Printf("总开支:%d",expense)
}

func main()  {
	se1 := seniorEmployee{1,3000,1000}
	se2 := seniorEmployee{2,3000,2000}
	e1 := Employee{3,3000}
	employees := []SalaryCalculate{se1,se2,e1}
	totalExpanse(employees)
}
