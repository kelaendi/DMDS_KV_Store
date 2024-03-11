package kvstore

import "testing"

func Test1(t *testing.T) {
	var kv = New()
	kv.Open()
	// value := new
	// kv.Put(1, value)
	kv.Get(1)
	kv.Close()
	kv.Delete()

	// KVStore.Open()

}
