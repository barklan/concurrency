---
layout: center
---

```go
func main() {
	var д1, д2, д3, д4 sync.Mutex
	перекресток := [4]*sync.Mutex{&д2, &д2, &д3, &д4}

	помеха_справа := func(i int) int {
		if i == 3 {
			return 0
		} else {
			return i + 1
		}
	}

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
