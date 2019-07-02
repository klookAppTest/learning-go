package main

import "fmt"

var ch chan int

func sum(s []int,c chan int){
	sum := 0
	for _,v := range s{
		sum += v
	}
	c <- sum
}

func count(num int){
	//fmt.Println(num)
	ch <- num
}

func main(){
	s := []int{7,2,8,-9,4,0}

	c := make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)

	x,y := <-c, <-c

	fmt.Println(x,y,x+y)

	ch = make(chan int)
	for i := 0;i<10;i++{
		go count(i)
	}


	for i := 0;i<10;i++{
		fmt.Println(<-ch)
	}

}