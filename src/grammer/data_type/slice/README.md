### array

#### basic info
- 长度是类型的组成部分
- 多维数组，len cap只返回第一维度
- 指针数组：元素是指针
- 数组指针：来用来修改数组的元素
- Go 数组是值类型，赋值和传参都会赋值整个数组数据


### slice
#### basic info
- 通过指针引用底层数组
- 设定相关属性，将数据读写操作限定在指定区域内
-----
- 创建切片无需先创建数组， make函数会自动完成底层数组内存分配

#### 两种创建方式 Test1
```go
var a []int // a==nil len=0, 不可读可append, 遍历不需要check nil

b := []int{} //b!=nil len=0

c := make([]int, 2, 10)//len=2 cap=10

d := make([]int, 2)//len=2 cap=2


```


#### append example(Test2)
```go
func main() {
	//当append向切片中追加元素时， 在空间足够的情况下， 新元素将从x[len(x)]位置开始存放， append会生成一个新的切片
	//但不会修改原切片x
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x) //1,2,3
	fmt.Println(y) //1,2,3,5
	fmt.Println(z) //1,2,3,5
}
```