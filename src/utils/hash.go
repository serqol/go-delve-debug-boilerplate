package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMapHash(data map[string]string) string {
	var dataString string
	for key, value := range data {
		dataString += key + value
	}
	return hash(dataString)
}

func hash(data string) string {
	bytes := md5.Sum([]byte(data))
	return hex.EncodeToString(bytes[:])
}
