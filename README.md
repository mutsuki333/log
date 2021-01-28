# log
A minimal go logger to extend with Info function

Mostly copied for the stdlib, so this package is fully compatible.   
All you need is to replace the import for the extended Info function.  
And this lib uses init() to override std log, so drop in replacement can be done without modifying existing code.


## Usage

```go
import "github.com/mutsuki333/log"

func main() {
    log.Println("This is for debug")
    log.Info("This is info")

    //Disable Debug log
    log.SetOutput(ioutil.Discard)
    log.Println("No more debug logs")
}
```

output:

> Debug: 2021/01/27 23:32:47 main.go:6: This is for debug  
> 2021/01/27 23:32:47 This is info
