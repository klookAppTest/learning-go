#### 常量
1.常量的定义格式  
```
const identifier [type] = value
```
>显示类型定义 const a string = 'abc'  
隐式类型定义 const a = 'abc'

2.常量枚举
```
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```

3.自增长
```
const (
    a = iota        // 0  
    b               // 1
    c               // 2
)
```

4.自定义类型
```
type output int
const (
    one output = iota       // 0
    two                     // 1
    three                   // 2
    four                    // 3
)
```

5.skipping value
>使用下划线跳过不想要的值
```
type output int
const (
    one output = iota       // 0
    two                     // 1
    three                   // 2
    _
    _
    five                    // 5
)
```

6.表达式
```
type ByteSize float64

const (
    _           = iota
    KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
    MB                             // 1 << (10*2)
    GB                             // 1 << (10*3)
    TB                             // 1 << (10*4)
    PB                             // 1 << (10*5)
    EB                             // 1 << (10*6)
    ZB                             // 1 << (10*7)
    YB                             // 1 << (10*8)
)
```

7.