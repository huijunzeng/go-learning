**lession2：**

Go 通过类型别名（alias types）和结构体的形式支持用户自定义类型，或者叫定制类型。  

```
type identifier struct {
    field1 type1
    field2 type2
    ...
}
```
组成结构体类型的那些数据称为 字段（fields）。每个字段都有一个类型和一个名字；在一个结构体中，字段名字必须是唯一的。  
结构体的概念在软件工程上旧的术语叫 ADT（抽象数据类型：Abstract Data Type），在一些老的编程语言中叫 记录（Record），比如 Cobol，在 C 家族的编程语言中它也存在，并且名字也是 struct，在面向对象的编程语言中，跟一个无方法的轻量级类一样。  
不过因为 Go 语言中没有类的概念，因此在 Go 中结构体有着更为重要的地位。  

使用 new 函数给一个新的结构体变量分配内存(跟java类似，包括给字段赋值以及取字段的值)，它返回指向已分配内存的指针：var t *T = new(T)，如果需要可以把这条语句放在不同的行（比如定义是包范围的，但是分配却没有必要在开始就做）。  


使用点号符可以获取结构体字段的值：structname.fieldname。  
在 Go 语言中这叫 选择器（selector）。无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的 选择器符（selector-notation） 来引用结构体的字段：
```
type myStruct struct { 
    i int 
    m string
    p float32
}

//func (t *type)printName() {} 表示定义type结构体的方法（相当于java中对象的方法） 可通过示例结构体直接调用
func (m *myStruct) printTest() error {
	fmt.Println("jack")
	return nil
}

var v myStruct    // v是结构体类型变量
var p *myStruct   // p是指向一个结构体类型变量的指针
v.i // 取变量值
p.i
v.printName // 调用结构体方法
```

初始化一个结构体实例（一个结构体字面量：struct-literal）的更简短和惯用的方式如下：
```
ms := &myStruct{10, "test", 15.5}
// 此时ms的类型是 *myStruct  &不是必须的 底层仍然会调用 new ()，这里值的顺序必须按照字段顺序来写  
// 假如是这种写法myStruct{i:10, m:"test", p:15.5}，则不必按顺序
```
或者：
```
var ms myStruct
ms = myStruct{10, "test", 15.5}
```