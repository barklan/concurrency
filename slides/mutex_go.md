---
layout: center
---

```go {all|2-3|13-15|all}
func main() {
	number := 0
	var m sync.Mutex

	num_of_routines := 1_000
	var wg sync.WaitGroup
	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		go func() {
			defer wg.Done()

			m.Lock()
			number++
			m.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(number)
}
```

<style>
code {
	font-size: 16px;
}
</style>
