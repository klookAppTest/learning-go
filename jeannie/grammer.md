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