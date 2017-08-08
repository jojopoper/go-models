package verify

import (
	"fmt"

	"github.com/dgryski/dgoogauth"
)

// GoolgeAuth google二次验证定义
type GoolgeAuth struct {
}

// VerifyTimebased 验证google Time base验证码是否正确
func (ths *GoolgeAuth) VerifyTimebased(key, code string, ws ...int) (bool, error) {
	var totp dgoogauth.OTPConfig
	totp.Secret = key
	if len(ws) == 0 {
		totp.WindowSize = 2
	} else {
		totp.WindowSize = ws[0]
	}
	b, err := totp.Authenticate(code)
	if err != nil {
		err = fmt.Errorf("verifyGoogleAuth has error : %v", err)
	}
	return b, err
}
