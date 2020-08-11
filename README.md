# 新版統一證號格式檢查器
針對2020年所執行的修正外來人口統一證號格式專案所釋出的驗證機制

# 如何使用
```go
package main
import (
    "fmt"
    "github.com/cychiang/tw-id-validator"
)

func main() {
    fmt.Println(validator.IdForeign("A800000014")) // true
    fmt.Println(validator.IdForeign("A800000010")) // false
}
```
