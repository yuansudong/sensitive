<h1 style="color:blue">
    1.	Description
</h1>

该项目主要用作于敏感词检测,并将检测到的敏感词屏蔽为*号.

<h1 style="color:blue">
    2.	Useage
</h1>

```shell
go get github.com/yuansudong/sensitive
```

<h1 style="color:blue">
    3.	Example
</h1>



```go
package main
import (
	"github.com/yuansudong/sensitive"
)

// _LoadSensitiveWords 加载敏感词
func _LoadSensitiveWords() []string {
    return []string{
        "haha",
        "hello",
        "悲",
    }
}


func main() {
    sensitive.Init(_LoadSensitiveWords)
    mInst :=  sensitive.Get()
    log.Println(mInst.Replace("haha,人生若如初见,何事秋风悲画扇"))
    mInst.Release()
}

```







