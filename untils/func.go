package untils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
	"strings"
	"bytes"
	"math/big"
	"crypto/rand"
)

type Encrypt struct{}

//将字符串加密成 md5
func (*Encrypt) EncodeMd5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return hex.EncodeToString(hash.Sum(nil))
}

//base64编码
func (*Encrypt) Base64Encode(raw []byte) string {
	t := base64.StdEncoding.EncodeToString(raw)
	t = strings.TrimSpace(t)
	t = strings.Replace(t, "\r", "", -1)
	t = strings.Replace(t, "\n", "", -1)
	t = strings.Replace(t, "\n\r", "", -1)
	t = strings.Replace(t, "\r\n", "", -1)
	return t
}

//base64解码
func (*Encrypt) Base64Decode(raw string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(raw)
}

//生成随机长度字符串
func RandomString(len int) string  {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0;i < len ;i++  {
		randomInt,_ := rand.Int(rand.Reader,bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}