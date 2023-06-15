package services

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

var chars = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")

func HashShortURL(URL string) string {
	hex := fmt.Sprintf("%x", md5.Sum([]byte(URL)))
	resURL := make([]string, 4)
	for i := 0; i < 4; i++ {
		val, _ := strconv.ParseInt(hex[i*8:i*8+8], 16, 0)
		lHexLong := val & 0x3fffffff
		outChars := ""
		for j := 0; j < 6; j++ {
			outChars += chars[0x0000003D&lHexLong]
			lHexLong >>= 5
		}
		resURL[i] = outChars
	}

	return resURL[0]
}
