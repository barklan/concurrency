---
layout: center
---

```go
func main() {

    go func() {
        fmt.Println("Hi, I'm goroutine!")
    }()

    fmt.Println("Bye!")
}
```

<style>
code {
    font-size: 16px;
}
</style>
