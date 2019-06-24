#### go指针

>取地址 &  
取值 *

```
// swap.go
func swap(a, b *int) {
	t := *a
	*a = *b
	*b =t
}

func Test() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}
```