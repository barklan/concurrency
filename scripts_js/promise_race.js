import fetch from 'node-fetch';

const f1 = fetch('https://mail.ru')
const f2 = fetch('https://yandex.ru')


Promise.race([f1, f2]).then(result => {
    console.log(result.url) // second
})
