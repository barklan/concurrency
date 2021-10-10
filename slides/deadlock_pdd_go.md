---
layout: center
---

<div class="grid grid-cols-[1fr,2fr] gap-x-4"><div>

```go
func помеха_справа(i int) int {
	if i == 3 {
		return 0
	} else {
		return i + 1
	}
}
```

</div><div>

```go {all|2|3|5|13|14|15|17-18|all}
func main() {
	var д1, д2, д3, д4 sync.Mutex
	перекресток := [4]*sync.Mutex{&д2, &д2, &д3, &д4}

	машины := 4
	var wg sync.WaitGroup
	wg.Add(машины)

	for i := 0; i < машины; i++ {
		go func(дорога int) {
			defer wg.Done()

			перекресток[дорога].Lock()
			time.Sleep(time.Second)
			перекресток[помеха_справа(дорога)].Lock()

			перекресток[дорога].Unlock()
			перекресток[помеха_справа(дорога)].Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("All cars passed!")
}
```

</div></div>

<style>
code {
    font-size: 14px ;
}
</style>
