package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(input string, salt string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input + salt))
	return hex.EncodeToString(hash.Sum(nil))
}
