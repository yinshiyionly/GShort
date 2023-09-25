package services

import (
	"crypto/md5"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"
	"strings"
)

var chars = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")

// GenerateTinyUrlByHash 通过 hash 生成短链接字符串
func GenerateTinyUrlByHash(url string) string {
	hex := fmt.Sprintf("%x", md5.Sum([]byte(url)))
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

// GenerateTinyUrlByPrimaryKey 通过 主键 生成短链接字符串
func GenerateTinyUrlByPrimaryKey(id string) string {
	// MongoDB _id 格式化
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalf("Transformer objectID failed: %v", err)
	}
	// 设置随机数种子
	//rand.NewSource(time.Now().UnixNano())
	// 生成随机数
	//randomNumber := rand.Int63()
	// 防止溢出
	newId := objectID.Timestamp().UnixMicro()
	fmt.Println(newId)
	indexAry := encode62(newId)
	return getString62(indexAry)
}

// encode62 转换成62进制
func encode62(id int64) []int64 {
	var indexAry []int64
	base := int64(len(chars))

	for id > 0 { // i < 0 时,说明已经除尽了,已经到最高位,数字位已经没有了
		remainder := id % base
		indexAry = append(indexAry, remainder)
		id = id / base
	}

	return indexAry
}

// getString62 输出字符串, 长度不一定为6
func getString62(indexAry []int64) string {
	result := ""
	for val := range indexAry {
		result = result + chars[val]
	}

	return reverseString(result[:6])
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
