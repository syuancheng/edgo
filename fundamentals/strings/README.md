# 字符串

Go 字符串是只读字节序列，通常保存 UTF-8 文本，但不保证内容一定是合法 UTF-8。

- `len(s)` 返回字节数，不是字符数。
- `s[i]` 得到一个字节。
- `for range` 按 UTF-8 解码并产生字节偏移和 `rune`。
- `[]byte(s)` 用于处理原始字节，`[]rune(s)` 适合按 Unicode 码点处理。
- 频繁拼接优先使用 `strings.Builder`。
