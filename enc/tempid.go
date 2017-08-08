package codeenc

// TempID 临时ID
type TempID struct {
}

// GetID 获取临时ID
func (ths *TempID) GetID(seed, key string) string {
	aes := &AESCode{}
	ret, err := aes.AESEncode(key, []byte(seed))
	if err != nil {
		return ""
	}
	return ret
}

// GetSeed 获得seed
func (ths *TempID) GetSeed(encode, key string) string {
	aes := &AESCode{}
	ret, err := aes.AESDecode(key, encode)
	if err != nil {
		return ""
	}
	return string(ret)
}
