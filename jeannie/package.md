#### go包
属于某一个包的源文件都应该放置于一个单独命名的文件夹里。按照 Go 的惯例，应该用包名命名该文件夹  

文件夹结构
```
bin
pkg
src  ---文件夹
    geometry  ---文件夹
        geometry.go   ---go可执行程序
        rectangle   ---文件夹
            rectprops.go   ---go程序
```

在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。mian函数可以使用被导出的函数与变量  

```
//rectprops.go

func Area(len, wid float64) float64 {
	ares := len * wid
	//fmt.Println("Area: %.2f", ares)
	return ares
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len*len) + (wid*wid))
	//fmt.Println("Diagonal %.2f", diagonal)
	return diagonal
}
```
init函数  
空白标识符
