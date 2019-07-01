package main

import "fmt"

func main(){
	var a = 10
	var b = 20
	//var ptr1 *int
	//var ptr2 *int
	//ptr1 = &a
	//ptr2 = &b
	fmt.Printf("传参之前a的值：%d,b的值：%d\n",a,b)
	// &代表传入参数的地址
	c, d := simpleSwap(a,b)
	fmt.Printf("不用指针传入参数，a的值:%d ,b的值:%d\n",c,d)
	swap(&a,&b)

	fmt.Printf("传参之后a的值：%d,b的值：%d\n",a,b)
}
// 交换变量的数据，传入指针则会改变参数的数据
func swap(a *int, b *int)  {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
// go语言可以做到传回多个参数，不是通过数组或者其它形式
func simpleSwap(a,b int) (int ,int){
	return b,a
}