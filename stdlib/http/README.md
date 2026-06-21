# HTTP 客户端

生产代码应复用 `http.Client`，配置超时，并始终关闭响应体。收到非 2xx 状态码时 `client.Do` 通常不会返回错误，调用方必须检查 `StatusCode`。

示例使用 `httptest.Server`，不依赖真实网络，因此可重复运行。更细粒度的超时可通过请求 context 或自定义 `http.Transport` 配置。
