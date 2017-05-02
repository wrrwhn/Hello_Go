
[TOC]

# GO
## 简介
- [Github_Go-zh](https://github.com/Go-zh/go)
- [Go 指南](https://go-tour-zh.appspot.com/welcome/1)
- [Github_Go](https://github.com/golang/go)
- [Go入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN/tree/master/eBook)

## 安装
- [Go-Download](https://golang.org/dl/)
- [Install](https://go-zh.org/doc/install)
- add *D:\server\go\1.8* to `path`
- 添加 `GOPATH` 和 `GOROOT`=`GOBIN` 对应目录

## 进度
- [the-way-to-go_ZH_CN](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/03.5.md)
- [godoc -http=:9999](http://localhost:9999/doc/code.html)

## 入门
### 语言特性与环境
- 发展目标
    + 类型安全和内存安全
    + 将静态语言的安全性和高效性与动态语言的易开发性进行有机结合
    + **对于网络通信、并发和并行编程的极佳支持**
        * 通过 goroutine  实现
        * 通过 channel 来实现各个 goroutine 之间的通信
        * 分段栈增长和 goroutine 在线程基础上多路复用技术
    + 构建速度极快
        * 包模型，严格依赖检查
    + 支持调用 C 语言编写的库文件
- 指导设计原则
    + 仅提供一两种方法来解决某个办法，保证易读
- 语言的特性
    + 让所有的东西都是显式的
    + 支持交叉编译
- 语言的用途
    + web 服务器、存储集群、巨型中央服务器
    + 实现复杂事件处理（CEP），提供少量并行支持，调试抽象和高性能
    + 不适实时性要求很高的场景
- 特性缺失
    + 不支持隐式转换
    + 不支持泛型
    + 不支持动态加载代码
    + 不支持动态链接库
    + 不支持断言
    + 不支持静态变量

### 平台与架构
- 适用操作系统编译器
    + Linux
    + FreeBSD
    + Mac OS X
- 版本
    + GC 原生编译器
        * 移植至 windows 平台
        * 基于 Plan 9 操作系统使用的 C 工具链
        * 编译器（C 实现）
            - 将源代码编译为项目代码（程序文本）
            - 使用非分代、无压缩和并行方式编译，较 gccgo 快，但不可使用 gcc 进行链接
        * 链接器（C 实现）
            - 将项目代码链接到可执行的二进制文件（机器代码）
    + gccgo 非原生编译器
        * Windows 可通过 MinGW 使用
    + 均为单通道
- 标记
    + 描述
        * 通过命令行设置可选参数来影响构建特殊的目标结果
    + 编译器相关
        * `-I` 针对包的目录搜索
        * `-d` 打印声明信息
        * `-e` 不限制错误打印的个数
        * `-f` 打印栈结构
        * `-h` 发生错误时进入恐慌（panic）状态
        * `-o` 指定输出文件名 
        * `-S` 打印产生的汇编代码
        * `-V` 打印编译器版本 
        * `-u` 禁止使用 unsafe 包中的代码
        * `-w` 打印归类后的语法解析树
        * `-x` 打印 lex tokens
- 文件扩展与包
    + `.go`
    + 创建目录时，文件夹名称的**空格**应用**下划线**或**一般符号**替换

### 环境变量
- `$GOROOT`
    - 表示 Go 在你的电脑上的安装位置，它的值一般都是 $HOME/go。
- `$GOARCH`
    - 表示目标机器的处理器架构，它的值可以是 386、amd64 或 arm。
- `$GOOS`
    - 表示目标机器的操作系统，它的值可以是 darwin、freebsd、linux 或 windows。
- `$GOBIN`
    - 表示编译器和链接器的安装位置，默认是 $GOROOT/bin，如果你使用的是 Go 1.0.3 及以后的版本，一般情况下你可以将它的值设置为空，Go 将会使用前面提到的默认值。
- `$GOHOSTOS`
    - 交叉编译时使用，用于指定目标机器操作系统，默认同 `$GOOS`
- `$GOHOSTARCH`
    - 交叉编译时使用，用于指定目标机器处理器架构，默认同 `$GOARCH`
- `$GOPATH`
    - 默认采用和 $GOROOT 一样的值，但从 Go 1.1 版本开始，你必须修改为其它路径。它可以包含多个包含 Go 语言源码文件、包文件和可执行文件的路径，而这些路径下又必须分别包含三个规定的目录：src、pkg 和 bin，这三个目录分别用于存放源码文件、包文件和可执行文件。
- `$GOARM`
    - 专门针对基于 arm 架构的处理器，它的值可以是 5 或 6，默认为 6。
- `$GOMAXPROCS`
    - 用于设置应用程序可使用的处理器个数与核数。

### 安装目录清单
- `/bin`：包含可执行文件，如编译器，Go 工具
- `/doc`：包含示例程序，代码工具，本地文档等
- `/lib`：包含文档模版
- `/misc`：包含与支持 Go 编辑器有关的配置文件以及 cgo 的示例
- `/os_arch`：包含标准库的包的对象文件（.a）
- `/src`：包含源代码构建脚本和标准库的包的**完整源代码**
- `/src/cmd`：包含 Go 和 C 的编译器和命令行脚本

### Runtime
- 特点
    + C 语言编写，1.5 后自举（编译器自行编译自己的编译器）
    + runtime 嵌入至每一个可执行文件中
- 功能
    + 内存分配
    + 垃圾回收
        * 标记-清除回收器
    + 栈处理
    + goruntime
    + channel
    + slice
    + map
    + reflection

## 编辑器、开发环境与工具
### 调试
+ 打印相关变量值
    * `print/println`
    * `fmt.Print/fmt.Println/fmt.Printf`
+ `fmt.Printf`
    * `%+v` 打印包括字段在内的实例的完整信息
    * `%#v` 打印包括字段和限定类型名称在内的实例的完整信息
    * `%T` 打印某个类型的完整说明
+ `panic`
    * 获取堆栈跳跃信息
+ `defer`
    * 跳跃代码执行过程

### 构建& 执行
+ [`gofmt`](http://golang.org/cmd/gofmt/)
    * 构建前自动进行格式化保存
    * 指令
        - `gofmt –w program.go`
            + 格式化源文件代码，并覆盖原始内容
        - `gofmt program.go`
            + 仅打印格式化后的结果
        - `gofmt -w *.go`
            + 格式化并重写所有 Go 源文件
        - `gofmt map1`
            + 格式化并重写 map1 目录及其子目录下的所有 Go 源文件
        - `gofmt -r '(a) -> a' –w *.go`
            + 用双引号括起来的替换规则实现代码
+ 构建错误
    * xxx行 `a declared and not used`
+ 构建成功
    * `Program exited with code 0`
+ [`go build`](http://wiki.jikexueyuan.com/project/go-command-tutorial/0.1.html)
    * 编译并安装自身包和依赖包
+ [`go install`](http://wiki.jikexueyuan.com/project/go-command-tutorial/0.2.html)
    * `go build` + 安装编译后结果文件至指定目录
+ `go get`
    * 下载更新指定的代码、依赖包后编译、安装
+ [`go clean`](http://wiki.jikexueyuan.com/project/go-command-tutorial/0.4.html)
    * 清除 `go build`|`go test -c`
    * 参数
        - `-i` 删除当前代码包安装时所产生的结果文件
        - `-r` 删除包括当前代码包的所有依赖包的上述目录和文件
+ [`go run`](http://wiki.jikexueyuan.com/project/go-command-tutorial/0.6.html)
    * 编译并运行命令源码文件 `main/ main()`
    * `go run fileName.go -param ~paramVal`
    * 参数
        - `-n` 打印相关命令而 __不执行__
        - `-x` 打印相关命令且 __执行__
        - `-work` 展示临时工作目录的位置
+ [`go test`](http://wiki.jikexueyuan.com/project/go-command-tutorial/0.7.html)
    * 参数
        - `-c` 仅生成运行测试的可执行文件，但 __不执行__
        - `-i` 安装运行测试所需依赖包
        - `-o` 指定运行测试的可执行文件名称

### 代码文档
+ `go doc`
    * 作用
        - 从 Go 程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档
        - 仅能获取 Go 安装目录下的注释内容
    * 示例
        - `go doc package[/subpackage | function | ]`
+ [`godoc`](http://golang.org/cmd/godoc/)
    * 查阅
        - 执行 `godoc -http=:6060`
        - 浏览器打开 `http://localhost:6060/`

###其它工具
+ `go install` 
    * 安装 Go 包的工具，主要用于安装非标准库的包文件，将源代码编译成对象文件
+ `go fix` 
    * 用于将你的 Go 代码从旧的发行版迁移到最新的发行版，它主要负责简单的、重复的、枯燥无味的修改工作
    * 像 API 等复杂的函数修改，工具则会给出文件名和代码行数的提示以便让开发人员快速定位并升级代码。
    * Go 在标准库就提供生成抽象语法树和通过抽象语法树对代码进行还原的功能。该工具会尝试更新当前目录下的所有 Go 源文件，并在完成代码更新后在控制台输出相关的文件名称。
+ `go test`
    * 一个轻量级的单元测试框架

### 性能
+ Go.speed = 1.2 * [C++ | C]= Scala.speed/ [2|3]= Java.speed/ [5~10]
+ Go.build = [Java| C++].build/ 5= Scala.build/ 10
+ Go.size> all.size
+ Go.memory = C++.memory = Scala.memory/ 2= Java.memory/ 5

### 与其它语言交互
+ 与 C 交互
    * [cgo](http://golang.org/cmd/cgo)
        - 提供了对 FFI（外部函数接口）的支持，能够使用 Go 代码安全地调用 C 语言库
        - 引用
        ```
        import "C"
        import "unsafe"
        ```
        - 释放内存
        ```
        cs := C.CString(s)
        defer C.free(unsafe.Pointer(cs))    // 建议于创建后释放
        ```
        - 构建
            + 使用 `GOFILES` 指定 Makefile 文件
            + 使用 `CGOFILES` 罗列需使用 cgo 编译的文件列表
+ 与 C++ 交互
    * SWIG
        - 特点
            + 支持在 Linux 系统下使用 Go 代码调用 C 或者 C++ 代码
            + SWIG 会产生 C 的存根函数
            + 库可使用 cgo 调用
            + 相关 Go 文件可自动生成
        - 缺点
            + **无法支持所有 C++ 库**，如无法解析 TObject.h

## 文件名、关键字与标识符
### 命名
- 文件
    + 小写字母+ 下划线组成
    + 不可包含空格或其它特殊字符
- 参数
    + `[a-Z|_][a-Z|_|0-9]*`
- 关键字
    ```
    | break    | default     | func   | interface | select |
    | case     | defer       | go     | map       | struct |
    | chan     | else        | goto   | package   | switch |
    | const    | fallthrough | if     | range     | type   |
    | continue | for         | import | return    | var    |
    ```

### 结构和要素
- 示例
    ``` go
    package main
    import "fmt"

    func main() {
        fmt.Println("hello, world")
    }
    ```
- 包
    + 同其它语言的类库、命名空间的概念
    + 名称限定为小写字母
    + 依赖关系决定构建顺序
        * A.go 依赖 B.go，而 B.go 又依赖 C.go
            - 编译 C.go, B.go, 然后是 A.go
            - 为了编译 A.go, 编译器读取的是 B.go 而不是 C.go
    + **包调整或重新编译，则引用该包的客户端程序必须全部重新编译**
    + 导入包等同于包含该包所有代码对象
        * 包中所有代码对象的标识符必须唯一
- 标准库
    + 位置为 `$GOROOT/pkg/$GOOS_$GOARCH/` 
    + 引用
    ```
    import "fmt"
    import "os"

    import "fmt"; import "os"

    import (
       "fmt"
       "os"
    )
    ```
- 可见性
    + `[A-Z][a-Z|_|0-9]*`
        * 可为外部包代码使用，同 public
- 函数
    + init()> main()
    + `{` 须与方法声明置于同行，__编译器强制规定__
    + 编译器自动补充分号，作为语句结束
    + 首字母大写则可为外部包调用
    + 示例
    ```
    func functionName(param1 type1, param2 type2) (ret1 type1, ret2 type2) {
       return ret1, ret2
    }
    ```
- 注释
    + `//`
    + `/* */`
    + 建议格式
    ```
    // 简要注释
    //
    // 另起一行后，详细注释
    ```
- 类型
    + 分类
        * 基本：int/ float/ bool/ string
        * 结构：struct/ array/ slice/ map/ channel
            - 默认值为 nil
        * 描述：interface
    + var 声明时自动初始化为该类型零值
- 一般结构
    + import
    + var
    + init()
    + main()
    + func() order by [call| name]
    + type
- 执行顺序
    + 顺序导入被 main 包引用的包
        * 递归导入包，其中包仅被导入一次
        * 以相反顺序初始化包中初始化常量和变量，并调用 init()
    + 调用 main() 执行程序
- 类型转换
    + __显式转换__
    + `valB:= typeB(valA)`

### 常量
- `const identifier [type] = value`
- 未指定类型常量会于使用时，根据环境推断所需具备类型
- `编译时可确定`，可使用`内置函数`
- 示例
```
const beef, two, c = "eat", 2, "veg"
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const (
    Monday, Tuesday, Wednesday = 1, 2, 3
    Thursday, Friday, Saturday = 4, 5, 6
)
const (
    a = iota    // 0，每次遇 const 则重置为 0
    b = iota    // 1，每次使用自动 +1
    c           // 2
)
```

### 变量
- `var identifier type`
- 示例
```

var a, b *int   // 均声明为指针
var (           // 常见于声明全局变量
    a int = 15
    b = false   // 自动推断类型
    c := 1      // 于函数体内声明局部变量
)
```
- 值类型
    + 概念
        * 类型的变量直接指向存在于内存中的值
        * 拷贝时于内存中将值进行拷贝
        * int/ float/ bool/ string/ __array__/ __struct__ 均属于值类型
        * 变量的值存储于 __栈__ 中
    + $val = val.内存地址
- 引用类型
    + 概念
        * 引用类型的变量 val 存储的为 val 值所在 内存地址|指针|第一个字所在位置
        * 指针/ slices/ maps/ channel
        * 存储于 __堆__ 中，拥有更大内存空间，便于进行垃圾回收
- 打印
    + Printf(format, obj...)
        * `%s` 字符串标识符
        * `%v` 默认输出格式标识符
    + Sprintf = Printf + returnStr
    + Print
        * fmt.Print("Hello:", 23) = Hello:_23
    + Println
- 初始化声明
    + `a := 50`
    + `a, b, c := 5, 7, "abc"` 
        * 并行赋值
    + `a, b = b, a`
        * 变量值交换
    + `_, b = 5, 7`
        * `_`为抛弃值，只写变量，常用于获取函数多余返回值
    + 仅于函数内可用
- init()
    + 初始化全局变量，执行程序前的数据检验或修复

### 基本类型和运算符
- bool
- int
    + int
        * 与操作系统相同位数（32/ 64）
        * 分类
            - int8（-128 -> 127）
            - int16（-32768 -> 32767）
            - int32（-2,147,483,648 -> 2,147,483,647）
            - int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
    + uint
        * 长度规则同 int
        * 分类
            - uint8（0 -> 255）
            - uint16（0 -> 65,535）
            - uint32（0 -> 4,294,967,295）
            - uint64（0 -> 18,446,744,073,709,551,615）
    + uintptr
        * 长度与指针大小
- float
    + float32（+- 1e-45 -> +- 3.4 * 1e38）
        * 精确到小数点后7位
    + float64（+- 5 * 1e-324 -> 107 * 1e308）
        * 精确到小数点后15位
        * 推荐使用，`math` 中函数通用类型
- complex
    + complex64
        * 格式为 `re + im I` = `5+ 10i`= `float + float i`= `complex(float, float)`
    + complex128
- 位运算
    + `&`
    + `|`
    + `^`
- 一元运算符
    + `^` 按位补足
        * `^2 = ^10 = m^x =  [-01|01] ^ 10 = -11`
        * x 为无符号值，则 m 为 1
        * x 为有符号值，则 m 为 -1
    + `<<` 位左移
        * `1 << 10 = 1024`
    + `>>` 位右称
    + `<<=` 位移并赋值
        * `a= 1, a<<= 10, a=== 1024`
- 逻辑运算符
    + `== != < <= > >=`
- 算术运算符
    + `+ - * / %`
    + `val/ 0 -> +Inf 无穷大`
    + `-= += *= /= %=`
    + `++ --`
- 优先级
     + `^ !`
     + `* / % << >> & &^`
     + `+ - | ^`
     + `== != < <= >= >`
     + `<-`
     + `&&`
     + `||`
 - 文本
     + `unicode.IsLetter(chars)`
         * 是否为字母
     + `unicode.IsDigit(chars)`
         * 是否为数字
     + `unicode.IsSpace(chars)`
         * 是否为空白符号
 
 ### 字符串标识符
 - 占用 1-4 字节
     + ASCII 时占用 1 字节
     + 其它字符占用 2-4 字节
- `str[idx]` 读取指定位置字节
- `s += str`
- 常用库
    + strings
        * HasPrefix
        * HasSuffix
        * Contains
        * Index
        * Count
        * Replace
        * ToUpper
        * ToLower
        * TrimSpace
        * TrimLeft
        * Fields
        * Join
    + unicode/utf8

### 时间和日期
- `time.Now()`
- `time.Add(time.Second)`
- `time.Sleep(time.Second)`
- `time.After(time.Second)`
- `time.NewTicker(time.Microsecond)`
- `time.AfterFunc(time.Second, tSubThread)`
- `tNow.After(tAfter)`


### 指针
- `&val` = 变量内存地址（32 位系统时为 4 字节，64 系统时为 8 字节）
- 初始化指针
    + `var intP *int`
    + 未分配时，默认值为 `nil`
- 常用缩写格式为 `ptr`
- 获取指针指向地址的值
    + `val = *(&val)`
- 适用场景
    + `传递参数`时，只传递指针（4/8字节），而不传递变量拷贝，减少大量内存的使用


## 控制结构
### if-else
- 结构
    ```
    if condition-A {            // { 需与判断条件一行
        // todo
    }
    } else if condition-B {     // }需要与 else-if 同行
        // todo
    }
    } else {

    }
    ```

### multi return
    ```
    iErrVal, err := strconv.Atoi(sErrVal)
    // 推荐在最接近异常的位置进行处理、封装
    if nil == err {
        return
        // return err
        // os.Exit(1)
    }
    ```

### switch
    ```
    switch result := calculate(); {     // { 需要与 switch 一致；switch 后可不需添加条件
        case result < 0:                // 终于可以用条件语句了！
        case result > 0:
            fallthrough                 // 继续执行下一个条件内语句
        default:
    }
    ```

### for
    ```
    str := "hello 够"
    for idx, rune := range str {    // 可读取文本等
        fmt.Printf("%-2d \t %d \t %U \t '%c' %X\n", idx, rune, rune, rune, []byte(string(rune)))
    }
    ```

### break & continue
- break
    + 用于 for 循环时，仅__跳出最内部结构__
    + 用于 switch 或 select 时，仅__跳过当前代码块__，执行后续代码
- continue
    + 仅可用于 for 循环

### label & goto
- label
    + 大小写敏感，但__推荐全大写__


## 函数
### 基准
- Go 为编译型语言，函数顺序与执行无关
- 遵循 DRY(Don't Repeat Yourself) 准则
- __不支持函数重载__
- 类型
    + 带名称函数
    + 匿名或 lambda 函数
    + 方法
- 格式
    ```
    func 函数名(实参-> 形参){                   // 函数调用时，拷贝并传递给函数

    }
    func 函数名(函数签名)
    type tFunc func(int, int) int                   // 类型声明
    func typeCheck(values ...interface{}) {         // 支持多类型、可变长参数
        for _, val := range values {
        }
    }
    ```
- 无参数函数称之为 `niladic` 函数

### defer
- 允许在函数返回前执行指定语句，类似 `finally` 语句块
- 作用
    + 关闭文件流
    + 解除锁定
    + 打印最终报告
    + 关闭数据库

### 内置函数
- close
- len
    + 返回类型值的长度
- cap
    + 返回类型值的最大容量
- new
    + 为 __值类型、自定义类型__ 分配内存
    + new(type)= 分配类型 T 的 __0__ 并返回地址
- make
    + 为 __内置引用类型__(切片、map、管道) 分配内存
    + make(type)= 返回 T __初始化__ 后的值
- copy/ append
    + 复制、连接切片
- panic/ recover
    + 错误处理
- print/ println
    + 底层打印函数，部署时建议使用 `fmt`
- complex/ real imag
    + 创建、操作得数

### 方法参数化
- 示例
    ```
    func callFuncArgs(i int, f func(int, int) int) int {
        return f(i, i)
    }
    ```

### 闭包/ 匿名函数
- 示例
    ```
    func() { fmt.Println("start...") }()
    funcCaller := func(iVal int) { fmt.Printf("closure.funcCaller(%d)\n", iVal) }
    ```
- 用途
    + 配合 defer 使用


## 容器
### 数据
- 同类型、长度固定的数据项序列
- 长度最大为 `2GB`
- 初始化
    + `var arr [ARRAY_LENGTH]int`
    + `sArr := [...]string{"yqj", "lmm"} == []string{"yqj", "lmm"}`
    + `[]string{3: "yao", 4: "lmm"} == [5]string{3: "yao", 4: "lmm"}`

### 切片
- 动态窗口，是对 __数组__ 连续片段的 __引用__，__长度可变__
- 指向底层数组，容量大于切片所定义容量，__仅于无切片指向时释放__
- 0<= len(slide) <= cap(slide)= len(array)
- slide= point(to array)+ len+ cap
- 创建
    + `var slice []type = array[start:end]= array[:]`
    + `slice:= make([]type, len[, cap])`
- reSlice
```
array:= [10]int{0,1,2,3,4,5,6,7,8,9}
slice:= array[5:7]  // 5,6
slice:= slice[0:4]  // 5,6,7,8
slice:= slice[4:4+1]    // 9
```

### bytes
- 可变长数组，用于读写未知长度 bytes
- 定义
    + `var buffer bytes.Buffer`
    + `var r *bytes.Buffer = new(bytes.Buffer)`
- 常用方法
    + `buffer.WriteString(string)`
    + `buffer.String()`

## Map
### 声明
- `var m map[keyType]valType`
    + 未初始化时， `nil== m`
    + keyType
        * 可使用 == 操作符类型
        * string/ int/ float
        * point/ interface
    + valType
        * ALL
- 初始化
    + `var mp= map[string]int{"one": 1, "two": 2}`
    + `var mp= __make__(map[string]int)`
- 若声明容量，增长受限时 map 自动`加 1`

## 包
### 标准库
- `unsafe`: 包含了一些打破 Go 语言“类型安全”的命令，一般的程序中不会被使用，可用在 C/C++ 程序的调用中。
- `syscall`|`os`|`os/exec`:  
    - `os`: 提供给我们一个平台无关性的操作系统功能接口，采用类Unix设计，隐藏了不同操作系统间差异，让不同的文件系统和操作系统对象表现一致。  
    - `os/exec`: 提供我们运行外部操作系统命令和程序的方式。  
    - `syscall`: 底层的外部包，提供了操作系统底层调用的基本接口。
- `archive/tar` 和 `/zip-compress`：压缩(解压缩)文件功能。
- `fmt`|`io`|`bufio`|`path/filepath`|`flag`:  
    - `fmt`: 提供了格式化输入输出功能。  
    - `io`: 提供了基本输入输出功能，大多数是围绕系统功能的封装。  
    - `bufio`: 缓冲输入输出功能的封装。  
    - `path/filepath`: 用来操作在当前系统中的目标文件名路径。  
    - `flag`: 对命令行参数的操作。　　
- `strings`|`strconv`|`unicode`|`regexp`|`bytes`:  
    - `strings`: 提供对字符串的操作。  
    - `strconv`: 提供将字符串转换为基础类型的功能。
    - `unicode`: 为 unicode 型的字符串提供特殊的功能。
    - `regexp`: 正则表达式功能。  
    - `bytes`: 提供对字符型分片的操作。  
    - `index/suffixarray`: 子字符串快速查询。
- `math`|`math/cmath`|`math/big`|`math/rand`|`sort`:  
    - `math`: 基本的数学函数。  
    - `math/cmath`: 对复数的操作。  
    - `math/rand`: 伪随机数生成。  
    - `sort`: 为数组排序和自定义集合。  
    - `math/big`: 大数的实现和计算。  　　
- `container`|`/list-ring-heap`: 实现对集合的操作。  
    - `list`: 双链表。
    - `ring`: 环形链表。
- `time`|`log`:  
    - `time`: 日期和时间的基本操作。  
    - `log`: 记录程序运行时产生的日志。
- `encoding/json`|`encoding/xml`|`text/template`:
    - `encoding/json`: 读取并解码和写入并编码 JSON 数据。  
    - `encoding/xml`:简单的 XML1.0 解析器。  
    - `text/template`:生成像 HTML 一样的数据与文本混合的数据驱动模板。  
- `net`|`net/http`|`html`:
    - `net`: 网络数据的基本操作。  
    - `http`: 提供了一个可扩展的 HTTP 服务器和基础客户端，解析 HTTP 请求和回复。  
    - `html`: HTML5 解析器。  
- `runtime`: Go 程序运行时的交互操作，例如垃圾回收和协程创建。  
- `reflect`: 实现通过程序运行时反射，让程序操作任意类型的变量。  

### regexp
- 常用
    + `regexp.Compile`
    + `regex.FindAllString(search, -1)`
    + `regex.ReplaceAllString(search, "$3$2$1")`

### sync
- `once`
    + 事件仅执行一次
- `Pool`
    + 特点
        * 临时对象池
        * 线程安全，自动伸缩
        * 每次 GC 时都会清空对象
    + 示例
        ```
        bufPool := sync.Pool{
            New: func() interface{} {
                return new(bytes.Buffer)
            },
        }
        ```
    + 参考
        * [临时对象池](http://www.01happy.com/golang-sync-pool/)

### package
- 目录
    ```
    └─Hello_Go
        └─hello
            ├─branch
            │  └─utils
            │     └─helloPackage.go
            ├─pack
            │  └─utils
            │     └─helloPackage.go
            └─utils
               └─helloPackage.go
    ```
- 事项
    + 包名，使用短小的不含有 _(下划线)的小写单词来为文件命名
    + 在使用第三方包的时候，当源码和.a均已安装的情况下，编译器链接的是源码。
    + 使用第三方包源码，实际链接了以该最新源码编译的临时目录下的.a文件而已。
    + import后面的最后一个元素应该是路径，就是目录，并非包名。
    + 别名指代的是引用路径下唯一的那个包

### 













## test
### 参考
- [Go怎么写测试用例](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/11.3.md)
- [go test](http://wiki.jikexueyuan.com/project/go-command-tutorial/0.7.html)

### 测试
- 示例
```
package utils       // 文件名称必须以 _test.go 结尾

import (
    "testing"       // 必输引入
)

// go test -v
func TestHelloTest_test(t *testing.T) {     // 固定 Test 前缀，`t *testing.T` 参数
    if HelloTest("test") {
        t.Log("pass")                       // Log 记录测试信息
    } else {
        t.Error("fail")                     // Error| Errorf| FailNow| Fatal| FatalIf 等标识测试不通过
    }
}
```
- 输出
    + 指令
        * `go test -v`
    + 控制台输出
        ```
        === RUN   TestHelloTest_test
        --- PASS: TestHelloTest_test (0.00s)
            helloTest_test.go:10: pass
        PASS
        ok      hello/utils 0.398s
        ```

### 压力测试
- 示例
    ```
    package utils

    import (
        "testing"
    )

    // go test -test.bench=".*"
    func Benchmark_HelloTest(b *testing.B) {        // 固定 `Benchmark_` 前缀, `b *testing.B` 参数
        b.StopTimer()                       // 暂停压力计时
        b.Log("init")                           // 中间用于完成初始化操作，如文件读取、数据库连接等
        b.StartTimer()                      // 重新开始压力计时
        for i := 0; i < b.N; i++ {
            HelloTest("test")
        }
    }

    ```
- 输出
    + 指令
        * `go test -v -test.bench=".*" `
    + 控制台输出
        ```
        Benchmark_HelloTest-8       500000000            3.52 ns/op
        --- BENCH: Benchmark_HelloTest-8
            helloTest_test.go:19: init
            helloTest_test.go:19: init
            helloTest_test.go:19: init
            helloTest_test.go:19: init
            helloTest_test.go:19: init
            helloTest_test.go:19: init
        PASS
        ok      hello/utils 2.575s
        ```