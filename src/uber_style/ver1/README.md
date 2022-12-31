### interface

#### 不需要interface指针(style1)
应该将接口作为值进行传递，在这样的传递过程中，实质上传递的底层数据仍然可以是指针。
接口实质上在底层用两个字段表示：
- 一个指向某些特定类型信息的指针。您可以将其视为"type"。
- 数据指针。如果存储的数据是指针，则直接存储。如果存储的数据是一个值，则存储指向该值的指针。
如果希望接口方法修改基础数据，则必须使用指针传递 (将对象指针赋值给接口变量,eg: f2)。

#### interface 合理性验证
https://github.com/xxjwxc/uber_go_guide_cn#interface-%E5%90%88%E7%90%86%E6%80%A7%E9%AA%8C%E8%AF%81

#### 接收器 (receiver) 与接口(style2)



#### 零值 Mutex 是有效的(style3)
零值 sync.Mutex 和 sync.RWMutex 是有效的。所以指向 mutex 的指针基本是不必要的。
如果你使用结构体指针，mutex 应该作为结构体的非指针字段。即使该结构体不被导出，也不要直接把 mutex 嵌入到结构体中。

#### 在边界处拷贝 Slices 和 Maps(style4)
slices 和 maps 包含了指向底层数据的指针，因此在需要复制它们时要特别注意。
https://github.com/xxjwxc/uber_go_guide_cn#%E5%9C%A8%E8%BE%B9%E7%95%8C%E5%A4%84%E6%8B%B7%E8%B4%9D-slices-%E5%92%8C-maps

#### tiny style
- 使用 defer 释放资源，诸如文件和锁。
- 枚举从 1 开始(style5)
- 在处理时间时始终使用 "time" 包，因为它有助于以更安全、更准确的方式处理这些不正确的假设。(style5)


