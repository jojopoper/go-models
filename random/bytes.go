package random

import (
	"crypto/rand"
	"fmt"
)

// RBInstance 随机字节唯一实例
var RBInstance *RBytes

// RBytes 随机字节定义
type RBytes struct {
}

func init() {
	RBInstance = new(RBytes)
}

// RandomBytes 随机N个字节
func (ths *RBytes) RandomBytes(length int) (k []byte, err error) {
	//rand Read
	k = make([]byte, length)
	if _, err = rand.Read(k); err != nil {
		err = fmt.Errorf("rand.Read() error : %v", err)
	}
	return
}
