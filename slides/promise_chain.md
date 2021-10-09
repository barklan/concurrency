---
layout: center
---

```js {all|2-4|all}
fetch('https://yesno.wtf/api')
  .then(getstatus)
  .then(json)
  .then(data => {
    console.log('Request succeeded with JSON response', data)
  })
  .catch(error => {
    console.log('Request failed', error)
  })
```

<style>
code {
    font-size: 16px;
}
</style>
