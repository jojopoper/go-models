package random

import (
	"crypto/rand"
	"fmt"
	"sync"
)

// RBInstance 随机字节唯一实例
var RBInstance *RBytes

// RBytes 随机字节定义
type RBytes struct {
	sync.Mutex
}

func init() {
	RBInstance = new(RBytes)
}

// RandomBytes 随机N个字节
func (ths *RBytes) RandomBytes(length int) (k []byte, err error) {
	ths.Lock()
	defer ths.Unlock()
	//rand Read
	k = make([]byte, length)
	if _, err = rand.Read(k); err != nil {
		err = fmt.Errorf("rand.Read() error : %v", err)
	}
	return
}
