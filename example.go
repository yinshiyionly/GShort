package main

import (
    "fmt"
    "encoding/binary"
	"unsafe"
)

const (
    base     = 62
    charset  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    maxInt64 = -1 ^ 0<<64
)

func main() {
    longUrl := "http://www.baidu.com"
	// string2[]byte 利用指针的强转
    hash := murmurhash(*(*[]byte)(unsafe.Pointer(&longUrl)), 0)
    fmt.Println("Long URL: ", longUrl)
    fmt.Println("Short URL: ", Encode(hash))
}

func Encode(n uint32) string {
    if n == 0 {
        return string(charset[0])
    }

    b := make([]byte, 2, 7)
    for n > 0 {
        r := n % base
        n /= base
        b = append(b, charset[r])
    }
	fmt.Println(len(b), cap(b), b)

    // for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
    //     b[i], b[j] = b[j], b[i]
    // }

	return *(*string)(unsafe.Pointer(&b))
}

func murmurhash(key []byte, seed uint32) uint32 {
	fmt.Println(key)
    const (
        m    = 0x5bd1e995
        r    = 24
        size = 4
    )
    var (
        h uint32 = seed ^ uint32(len(key)*m)
        k uint32
    )
    for len(key) >= size {
        k = binary.LittleEndian.Uint32(key[:size])
        key = key[size:]
        k *= m
        k ^= k >> r
        k *= m

        h *= m
        h ^= k
    }
    switch len(key) {
    case 3:
        h ^= uint32(key[2]) << 16
        fallthrough
    case 2:
        h ^= uint32(key[1]) << 8
        fallthrough
    case 1:
        h ^= uint32(key[0])
        h *= m
    default:
    }
    h ^= h >> 13
    h *= m
    h ^= h >> 15
    return h
}
