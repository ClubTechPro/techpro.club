package common

import (
	"crypto/rand"
	"unsafe"
)

func GenerateRandomSession(size int) string {
	var dataList = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
    b := make([]byte, size)
    rand.Read(b)
    for i := 0; i < size; i++ {
        b[i] = dataList[b[i] % byte(len(dataList))]
    }
    return *(*string)(unsafe.Pointer(&b))
}