package com

func Try(userFn func(), catchFn func(err interface{}), finalFn func()) {
    try(userFn, catchFn)
    finalFn()
}

func try(userFn func(), catchFn func(err interface{})) {
    defer func() {
        if err := recover(); err != nil {
            catchFn(err)
        }
    }()
    userFn()
}
