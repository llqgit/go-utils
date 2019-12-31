# Golang Time Utils

## include functions

### time
- [x] WeeksInYear(now time.Time) (weeks, year int)
- [x] GetRandomNum() int64

### 获取一个随机数
```go
GetRandomNum(max int64) int64

// eg
num := GetRandomNum(666)
```
### WeeksInYear 计算时间是当年的第几周
周一为每周第一天，根据本周周一判断本周是哪年的

```go
weeks, year := WeeksInYear(time.Now())
```