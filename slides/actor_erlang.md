---
layout: center
---

<div>

<span style="color:grey">Erlang</span>

```erlang {all|9-13|10-12|16-18|17|4-7|all}
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

</div>

<style>
code {
    font-size: 16px ;
}
</style>
