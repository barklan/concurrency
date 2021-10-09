---
layout: center
---

<div class="grid grid-cols-2 gap-x-4"><div>

```js
function greeting() {
    console.log('Hi Glebushek!')
}

setTimeout(greeting, 500)



console.log('Bye!');
```

</div>

<div v-click>

```js




setTimeout(function greeting() {
    console.log('Hi Glebushek!')
}, 500)

console.log('Bye!');
```

</div></div>

<style>
code {
    font-size: 16px ;
}
</style>
