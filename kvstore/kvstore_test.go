package kvstore

import (
	"testing"
)

func TestPutGet(t *testing.T) {
	controller := NewKVStoreControl()
	kv := controller.NewKVStore()
	var givenValue [10]byte
	givenValue[0] = 123
	kv.Put(1, givenValue)
	gottenValue := kv.Get(1)
	if gottenValue[0] != 123 {
		t.Errorf("value %v is not %v", gottenValue, givenValue)
	}
}

func TestOpenClose(t *testing.T) {
	controller := NewKVStoreControl()
	kv := controller.NewKVStore()
	errOpen := controller.Open(kv)
	if errOpen != nil {
		t.Errorf(errOpen.Error())
	}
	errClose := controller.Close(kv)
	if errClose != nil {
		t.Errorf(errClose.Error())
	}
}

func TestDeleteKV(t *testing.T) {
	controller := NewKVStoreControl()
	kv := controller.NewKVStore()
	errDeleteKV := controller.Delete(kv)
	if errDeleteKV != nil {
		t.Errorf(errDeleteKV.Error())
	}
}
