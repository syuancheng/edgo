# Gzip 解压

示例通过 `embed` 打包测试数据，避免硬编码本机绝对路径。`gzip.Reader` 是需要关闭的资源；对不可信输入还应使用 `io.LimitReader` 限制解压后大小，防止压缩炸弹耗尽内存。
