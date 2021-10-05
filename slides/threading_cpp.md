---
layout: center
---

# Threading

C++, Java, Rust, Haskell,...

```cpp {all|10|4-6|14|all}
#include <iostream>
#include <thread>

void foo(int a) {
    std::cout << a << '\n';
}

int main() {

    std::thread thread(foo, 10);

    // Keep going; the thread is executed separately

    thread.join();

    return 0;
}
```

<style>
code {
    font-size: 16px ;
}
</style>
