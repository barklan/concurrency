---
layout: full
---

...but,

<h2>C<span class="text-red-500">Python</span> has GIL</h2>

<br>

```c
static PyThread_type_lock interpreter_lock = 0; /* This is the GIL */
```

<span style="font-size: 8px; bottom: 50px; position: fixed; color: grey;">...which is not a problem exactly - but for wrong reasons</span>

<br>
<br>

<h2><span class="text-red-500">Ruby</span> MRI has GVL</h2>
