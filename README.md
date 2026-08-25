# go-cronexpr

Cron 表达式引擎：AST 解析 + 字段位图求值 + 时区层。配套 planner/ Web 规划台（:8222）。

```text
go build ./...
go test ./... -count=1
go run ./cmd/crond -addr :8222
```
