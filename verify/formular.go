package verify

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

// EFormular 简单加减验算
type EFormular struct {
	Value1   string `json:"value1"`
	Value2   string `json:"value2"`
	Value3   string `json:"value3"`
	Operator string `json:"operator"`
	ErrorMsg error  `json:"error"`
	ViCode   string
}

// New 新建一个计算公式验证内容
func (ths *EFormular) New() *EFormular {
	ret, err := ths.random(3, 30)
	if err == nil {
		if ret[0] > 15 {
			ths.Operator = "-"
		} else {
			ths.Operator = "+"
		}

		if ret[0] > 20 {
			ths.Value1 = fmt.Sprintf("%d", ths.max(int(ret[1]), int(ret[2])))
			ths.Value2 = fmt.Sprintf("%d", ths.min(int(ret[1]), int(ret[2])))
			ths.Value3 = "?"
		} else if ret[0] > 10 {
			ths.Value1 = fmt.Sprintf("%d", ths.max(int(ret[1]), int(ret[2])))
			ths.Value2 = "?"
			ths.Value3 = fmt.Sprintf("%d", ths.min(int(ret[1]), int(ret[2])))
		} else {
			ths.Value1 = "?"
			ths.Value2 = fmt.Sprintf("%d", ths.min(int(ret[1]), int(ret[2])))
			ths.Value3 = fmt.Sprintf("%d", ths.max(int(ret[1]), int(ret[2])))
		}
	}

	ths.ErrorMsg = err
	return ths
}

// Verify 验证一个公式
func (ths *EFormular) Verify() bool {
	if len(ths.ViCode) < 1 {
		return false
	}

	var v1, v2, v3 int

	if ths.Value1 == "?" {
		v1, _ = strconv.Atoi(ths.ViCode)
	} else {
		v1, _ = strconv.Atoi(ths.Value1)
	}

	if ths.Value2 == "?" {
		v2, _ = strconv.Atoi(ths.ViCode)
	} else {
		v2, _ = strconv.Atoi(ths.Value2)
	}

	if ths.Value3 == "?" {
		v3, _ = strconv.Atoi(ths.ViCode)
	} else {
		v3, _ = strconv.Atoi(ths.Value3)
	}

	if ths.Operator == "-" {
		return (v1 - v2) == v3
	}
	return (v1 + v2) == v3
}

func (ths *EFormular) random(length, max int) ([]byte, error) {
	//rand Read
	k := make([]byte, length)
	if _, err := rand.Read(k); err != nil {
		return nil, err
	}
	for i := 0; i < length; i++ {
		k[i] = k[i] % byte(max+1)
	}
	return k, nil
}

func (ths *EFormular) max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func (ths *EFormular) min(val1, val2 int) int {
	if val1 > val2 {
		return val2
	}
	return val1
}
