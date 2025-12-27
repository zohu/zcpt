package zcpt

import "testing"

func TestSha256(t *testing.T) {
	t.Log(Sha256("654321"))
	t.Log(Sha256(123456))
	t.Log(Sha256(nil))

	var map1 = map[string]string{"a": "1", "b": "2"}
	var map2 = map[string]string{"b": "2", "a": "1"}
	t.Log(Sha256(map1))
	t.Log(Sha256(map2))
}
