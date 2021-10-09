const bar = () => {
    throw new Exception()
}

const baz = () => console.log('baz')

const foo = () => {
    console.log('foo')
    bar()
    baz()
}

foo()
