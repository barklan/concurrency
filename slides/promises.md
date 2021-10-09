---
layout: center
---

```js
done = true

const myPromise = new Promise((resolve, reject) => {
    if (done) {
        resolve("Done!")
    } else {
        reject("Not done yet..")
    }
})
```

<br>

<v-click>

```js
function checkIfDone() {
    myPromise
        .then(ok => { console.log(ok) })
        .catch(err => { console.log(err) })
}

checkIfDone()
```

</v-click>

<style>
code {
    font-size: 16px ;
}
</style>
