package zcpt

import (
	"crypto/md5"
	"fmt"

	"github.com/bytedance/sonic"
)

func Md5(v any) string {
	if IsNil(v) {
		return ""
	}
	d, _ := sonic.Marshal(v)
	has := md5.Sum(d)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
