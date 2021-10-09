---
layout: center
---

<div class="grid grid-cols-2 gap-x-4"><div>

Before

```js
const getFirstUserData = () => {
  return fetch('/users.json')
    .then(response => response.json())
    .then(users => users[0])
    .then(user => fetch(`/users/${user.name}`))
    .then(userResponse => userResponse.json())
}


getFirstUserData()
```

</div>

<div v-click>

After

```js
const getFirstUserData = async () => {
  const response = await fetch('/users.json')
  const users = await response.json()
  const user = users[0]
  const userResponse = await fetch(`/users/${user.name}`)
  const userData = await userResponse.json()
  return userData
}

getFirstUserData()
```

</div></div>

<style>
code {
    font-size: 12px ;
}
</style>
