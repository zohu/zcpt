package zcpt

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"reflect"

	"github.com/bytedance/sonic/encoder"
)

func Sha256(v any) string {
	if IsNil(v) {
		return ""
	}
	d, _ := encoder.Encode(v, encoder.SortMapKeys)
	hash := sha256.Sum256(d)
	return hex.EncodeToString(hash[:])
}
func Sha256Stream(r io.Reader) string {
	hasher := sha256.New()
	_, err := io.Copy(hasher, r)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return true
	}
	switch rv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	default:
		return false
	}
}
