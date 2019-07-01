#### go语法

1.如果函数最后一个参数被记作 ...T ，这时函数可以接受任意个 T 类型参数作为最后一个参数

2.可变参数转化为切片
```
nums ...int
int []int{}
```

3.追加切片元素：当新的元素被添加到切片时，会创建一个新的数组，现有数组的元素被复制到这个新数组，并返回新数组的新切片引用，append改变的是副本（新数组）的len\cap，原slice（数组）的len\cap不变

4.map
// 创建
```
var countryMap map[string]string
countryMap = make(map[string]string)
countryMap["China"] = "Beijing"
...
or
countryMap := map[string]string {
    countryMap["China"] = "Beijing"
    ...
}

// 遍历
for key, value := range countryMap {
    ...
}

// 删除
delete(countryMap, "China")

// 长度
len(countryMap)
```

5.Go语言中的循环语句只支持for关键字，而不支持while和do-while结构。
```
// 无限循环
var op int
fmt.Scanln(&op)
for {
    if op == 0 {
        break //退出无限循环
    }
    // do something...
}
```

6.break continue goto
```
var a int = 10
LOOP: for a < 20 {
    if a == 15 {
        a = a + 1
        goto LOOP
    }
    // do something...
}
```

7.switch
> fallthrough
```
switch op {
		case 1:
			PersonSalary()
			fallthrough
		case 2:
			FoundPersonSalary()
		case 3:
			RangePersonSalary()
		default:
			fmt.Println("DEFAULT...")
		}
```

8.Scanln Scanf Sscan函数 bufio
- Scanln 带空格的字符串值保存到变量
- Scanf 指定参数
- Sscan系列函数 从字符串变量里读取
- bufio 缓存

9.结构体
> 如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。同样，如果结构体里的字段首字母大写，它也能被其他包访问到。
```
// 导出结构体
type Spec struct {
    Maker string //导出
    model string //未导出
    Price int // 导出
}
```
