package main

import (
	"github.com/kelaendi/DMDS_KV_Store/kvstore"
)

func main() {
	kv := kvstore.NewKVStore()
	var givenValue [10]byte
	givenValue[0] = 123
	kv.Put(1, givenValue)
}
