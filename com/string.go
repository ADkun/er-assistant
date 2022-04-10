package com

import (
    "strings"
    "strconv"
    "regexp"
    "time"
    "math/rand"
    "fmt"
)

func Trim(s string) string {
    res := strings.Replace(s, "\n", "", -1)
    res = strings.Replace(res, "\r", "", -1)
    res = strings.Trim(res, " ")
    return res
}

func GetRandStr(length int) string {
    if length < 1 {
        return ""
    }
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    charArr := strings.Split(char, "")
    charLen := len(charArr)
    ran := rand.New(rand.NewSource(time.Now().Unix()))
    var rchar string = ""
    for i := 1; i <= length; i++ {
        rchar += charArr[ran.Intn(charLen)]
    }
    return rchar
}

func IsDigit(s string) bool {
    pat := "\\d+"
    res, err := regexp.MatchString(pat, s)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("regexp.MatchString(%s, %s)执行失败", pat, s), err)
    }
    return res
}

func A2I(s string) int {
    res, err := strconv.Atoi(s)
    if err != nil {
        PanicErr(FuncName(), fmt.Sprintf("strconv.Atoi(%s)执行失败", s), err)
    }
    return res
}

func I2A(i int) string {
    return strconv.Itoa(i)
}

func GetCurTimeStamp() string {
    milli := time.Now().UnixNano() / 1e6
    return strconv.FormatInt(milli, 10)
}
