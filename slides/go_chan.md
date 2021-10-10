---
layout: image-right
image: /img/channel.jpg
---

<br><br><br><br>

```go{1|3-5|7|all}
c := make(chan int)

go func() {
    c <- 4
}()

value := <- c
```

<style>
code {
    font-size: 16px ;
}
</style>
