# HTTP 请求草稿

`.http` 文件可由常见编辑器的 HTTP Client 手动发送，不参与 Go 构建。通过变量配置本地地址和开发凭据，不要提交真实 token、cookie、内部域名或个人地址。可重复的 HTTP 行为测试应优先使用 [`httptest`](../../standard-library/http/)。
