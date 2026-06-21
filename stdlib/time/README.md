# Time

使用 `time.Time` 和 `time.Duration` 表达时间，避免用无单位整数。`Time.Truncate` 按绝对时长截断，适合固定间隔；日历日期运算使用 `AddDate`，时区使用 `time.Location`。

跨系统传输时明确时区和格式。比较时间优先使用 `Equal`、`Before` 和 `After`，不要依赖包含单调时钟数据的结构体直接比较。
