# Struct

### basic info
#### 基本规则
- 字段名唯一
- 可用_补位
- 可用自身指针
- 字段读顺序也是类型组成部分

#### Test1 
##### 匿名结构
```go
u := User{
	name: "tom",
	age: 12
}

u := struct{
	name string
	age int
}{
	name: "tom",
	age: 23
}
```

##### 比较 
- 只有struct中所有字段都支持==的时候

#### Test2
##### 指针
一级直至可以修改字段，多级不可以

#### Test3
##### struct{}

#### Test4
##### 匿名字段
也叫嵌入字段或嵌入类型
对于类型为结构体的字段，显示指定时与其他类型没有区别，仅仅代表某种类型的字段。而隐式指定时，原结构体的**字段**和**方法**看起来就像是被**继承**过来一样。
如下，Cat中嵌入了另一个结构体Animal，此时结构体Animal中的字段和方法会被提升到Cat中，看上去就像是Cat的原生字段和方法。
```go
type Animal struct {
	Name string
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

type Cat struct {
	Animal
}

type Dog struct {
	a Animal
}
```

#### Test5
标签

#### Test6
结构体中字段首字母小写时，包外不可见。在当前包中大小写可见性一致。