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

#### Test5
标签

#### Test6
结构体中字段首字母小写时，包外不可见。在当前包中大小写可见性一致。