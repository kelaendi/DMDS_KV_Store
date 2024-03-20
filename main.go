package main

import (
	"github.com/kelaendi/DMDS_KV_Store/kvstore"
)

func main() {
	controller := kvstore.NewKVStoreControl()
	kv := controller.NewKVStore()
	var givenValue [10]byte
	givenValue[0] = 123
	kv.Put(1, givenValue)
}
