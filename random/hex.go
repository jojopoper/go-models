package random

import "encoding/hex"

// RHInstance 随机16进制字符串唯一实例
var RHInstance *RHex

// RHex 随机16进制字符串
type RHex struct {
	RBytes
}

func init() {
	RHInstance = new(RHex)
}

// RandomHexStr 随机16进制字符串
func (ths *RHex) RandomHexStr(length int) (ret string, err error) {
	k, err := ths.RandomBytes(length)
	if err == nil {
		ret = hex.EncodeToString(k)
	}
	return
}
