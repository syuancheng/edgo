# Chapter 5 数据

## 字符串

## 数组

## 切片

## 字典

## 结构
### what is struct?
- 将多个不同类型命名字段序列打包成一个符合类型。
- 要求：
  - 字段名唯一
  - 可用`_`补位
  - 可用自身指针
  - 字段名排列顺序也是类型组成部分
  - 结构体中字段首字母小写时，包外不可见。在当前包中大小写可见性一致
- 初始化
  - 按顺序
  ```go
  type user struct {
    name string
    age byte
  }
  u1 := user{"Tom", 12}
  u2 := user{"Krystal"}//error: too few values in struct initializer
  ```
  - 命名（推荐使用命名初始化， 这样在扩充结构字段或调整字段顺序时，不会导致初始化语句出错）
  ```go
  u := User{
	name: "tom",
	age: 12
  }
  ```
- 匿名结构
```go
u := struct{
	name string
	age int
}{
	name: "tom",
	age: 23
}

type file struct {
    name string
    attr struct {
        owner int
        perm int
    }
}

```
- 比较: 只有struct中所有字段都支持`==`的时候 [ex_5_5_1.go](./ex_5_5_1.go)
- 指针：一级指针可以修改字段，多级不可以 [ex_5_5_2.go](./ex_5_5_2.go)

### 空结构
- `struct{}` 没有字段的结构类型 [ex_5_5_3.go](./ex_5_5_3.go)
  - 没有分配数组内存，但依然可以操作元素，对应切片的len cap属性也正常
  - 空接口可以用作通道元素类型，用于事件通知
### 匿名字段 [ex_5_5_4](./ex_5_5_4.go)
- def:没有名字，只有类型的字段，也叫`嵌入字段`或`嵌入类型`
- 哪些类型可以做
  - 结构体
  - 除了接口指针和多级指针以外的任何命名类型都可以
  - 不能将基础类型和其指针同时嵌入，因为两者隐式名字相同
- 显示和隐式区别
  - 显示指定时与其他类型没有区别，仅仅代表某种类型的字段。而
  - 隐式指定时，原结构体的**字段**和**方法**看起来就像是被**继承**过来一样。
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

### 字段标签 [ex_5_5_5](./ex_5_5_5.go)
- 字段标签并不是注释，而是用来对字段进行描述的元数据。
- 在运行期，可以用反射获取标签信息。