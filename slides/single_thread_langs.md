---
layout: full
---

...but,

<h2>C<span class="text-red-500">Python</span> has GIL</h2>

<br>

```c
static PyThread_type_lock interpreter_lock = 0; /* This is the GIL */
```

<br>
<br>

<h2 v-click><span class="text-red-500">Ruby</span> MRI has GVL</h2>

<br>
<br>

<h2 v-click>Node.<span class="text-red-500">js</span> (based on V8) is single threaded by design</h2>
