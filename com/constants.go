package com

import (
    "runtime"
)

var SLASH string
func init() {
    if runtime.GOOS == "linux" {
        SLASH = "/"
    } else {
        SLASH = "\\"
    }
}
