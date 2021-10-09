done = true

const myPromise = new Promise((resolve, reject) => {
    if (done) {
        resolve("Done!")
    } else {
        reject("Not done yet..")
    }
})

function checkIfDone() {
    myPromise
        .then(ok => { console.log(ok) })
        .catch(err => { console.log(err) })
}

checkIfDone()
