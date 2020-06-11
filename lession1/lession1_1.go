package main // 必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包

import "fmt" // 引入需要的依赖或包

const value int = 5

var name string = "狗娃"
var a = "G"
const (
	Unknown = 0
	Female = 1
	Male = 2
)

func main()  { // 函数体
	defer println("world") // 相当于java中的try{}finally{}语法体的finally  程序返回数据体后再执行
	name := "狗蛋"
	//name := "狗蛋"  只能使用 := 对同一个变量初始化声明
	name = "狗子" //但可以通过 = 赋值多次，相当于赋予一个新的值
	fmt.Println(name, ", hello") // 打印语句

	var c = add(1,2)
	fmt.Println("add_sum: ", c) // 打印语句
}

func add(a, b int) int {
	return (a + b + value)
}
