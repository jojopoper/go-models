package random

import (
	"encoding/base32"
)

// RB32Instance 随机Base32字符唯一实例
var RB32Instance *RBase32

// RBase32 随机Base32字符定义
type RBase32 struct {
	RBytes
}

func init() {
	RB32Instance = new(RBase32)
}

// RandomBase32 随机N个Base32字符
func (ths *RBase32) RandomBase32(length int) (ret string, err error) {
	rd, err := ths.RandomBytes(length)
	if err == nil {
		ret = base32.StdEncoding.EncodeToString(rd)
	}
	return
}
