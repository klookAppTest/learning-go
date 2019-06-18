# learning-go


<img src="./logo.gif"  height="150">


### 开始

1、安装go 环境

https://studygolang.com/dl

2、设置 ```GOPATH```, ```GOROOT``` 的环境变量。

```shell
$ go env

GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/tech/Library/Caches/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/tech/go"
GOPROXY=""
GORACE=""
GOROOT="/usr/local/Cellar/go/1.11.5/libexec"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.11.5/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/44/78c0cnzn0jn132wjcjlyq4d40000gn/T/go-build729519467=/tmp/go-build -gno-record-gcc-switches -fno-common"

```

4、 安装编辑器或IDE， 推荐 ```VS code``` 或 ```IntelliJ Go```。

5、 第一个go程序，```hello.go```

```go
package main

import "fmt"

func main(){
    fmt.Println("hello, go! ")
}
```

6、 运行程序

```shell
$ go run hello.go

hello, go!
```