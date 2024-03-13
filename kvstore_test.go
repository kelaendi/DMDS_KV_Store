package kvstore

import (
	"testing"
)

func TestPutGet(t *testing.T) {
	kv := NewKVStore()
	var givenValue [10]byte
	givenValue[0] = 123
	kv.Put(1, givenValue)
	gottenValue := kv.Get(1)
	if gottenValue[0] != 123 {
		t.Errorf("value %v is not %v", gottenValue, givenValue)
	}
}

func TestOpenClose(t *testing.T) {
	kv := NewKVStore()
	errOpen := kv.Open()
	if errOpen != nil {
		t.Errorf(errOpen.Error())
	}
	errClose := kv.Close()
	if errClose != nil {
		t.Errorf(errClose.Error())
	}
	kv.Close()
}

func TestDeleteKV(t *testing.T) {
	kv := NewKVStore()
	errDeleteKV := kv.Delete()
	if errDeleteKV != nil {
		t.Errorf(errDeleteKV.Error())
	}
}
