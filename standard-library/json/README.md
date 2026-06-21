# JSON

`encoding/json` 只处理导出字段。结构体标签控制字段名和省略规则；解码时必须传入可写指针。

处理外部输入时应检查所有错误。需要拒绝未知字段时使用 `Decoder.DisallowUnknownFields`；处理连续 JSON 值或流式输入时使用 `Decoder`，处理单个内存值时可使用 `json.Unmarshal`。
