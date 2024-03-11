package kvstore

import "testing"

func Test1(t *testing.T) {
	var kv = NewKVStore()
	kv.Open()
	// value := new
	// kv.Put(1, value)
	value := kv.Get(1)
	t.Errorf("value: %v", value)

	kv.Close()
	kv.Delete()

	// KVStore.Open()

}
