---
layout: center
---

```go {all|2-3|13-15|all}
func main() {
	channel := make(chan int, 1)
	channel <- 0

	num_of_routines := 1_000
	var wg sync.WaitGroup
	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		go func() {
			defer wg.Done()

			local_number := <-channel
			time.Sleep(1 * time.Microsecond)
			channel <- local_number + 1
		}()
	}

	wg.Wait()
	fmt.Println(<-channel)
}
```

<style>
code {
	font-size: 16px;
}
</style>
