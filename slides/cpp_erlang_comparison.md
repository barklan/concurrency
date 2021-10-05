<div class="grid grid-cols-2 gap-x-4"><div>

```cpp
#include <iostream>
#include <thread>

void foo(int a) {
    std::cout << a << '\n';
}

int main() {
    std::thread thread(foo, 10);
    thread.join();
    return 0;
}
```

</div><div v-click>


```erlang
-module(counter).
-export([run/0, counter/1]).

run() ->
    S = spawn(counter, counter, [0]),
    send_msgs(S, 100000),
    S.

counter(Sum) ->
    receive
        value -> io:fwrite("Value is ~w~n", [Sum]);
        {inc, Amount} -> counter(Sum+Amount)
    end.

send_msgs(_, 0) -> true;
send_msgs(S, Count) ->
    S ! {inc, 1},
    send_msgs(S, Count-1).
```

</div></div>
