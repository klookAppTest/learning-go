####  go信道
chan T 表示 T 类型的信道

简介有效的赋值
```
a := make(chan int)
data := <- a // 读取信道 a  
a <- data // 写入信道 a
```

```
package main

import "fmt"

func calcSquares(number int, Squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit *digit
		number /= 10
	}
	Squareop <- sum
}


func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit  := number % 10
		sum = digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}


func main() {
	number := 539
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <- sqrch, <-cubech
	fmt.Println("Final output %d\n", squares + cubes)
}

```