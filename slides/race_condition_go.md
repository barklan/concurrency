---
layout: full
---

```go
// multicore race -> one core false safety
// -> breaking the safety with timer
func main() {
	// runtime.GOMAXPROCS(1)

	number := 0

	num_of_routines := 1_000
	var wg sync.WaitGroup
	wg.Add(num_of_routines)

	for i := 0; i < num_of_routines; i++ {
		go func() {
			defer wg.Done()

			local_number := number
			time.Sleep(time.Microsecond * 1)
			number = local_number + 1
		}()
	}

	wg.Wait()
	fmt.Println(number)
}
```
