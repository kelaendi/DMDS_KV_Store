package kvstore

type KVStore interface {
	Open() error
	Close() error
	Delete() error
	Put(key uint64, value [10]byte)
	Get(key uint64) (value [10]byte)
}

type Option func(kv KVStore) KVStore

func NewKVStore(options ...Option) (kv KVStore) {
	return nil
	// for _, o := range options {
	// 	kv = o(kv)
	// }
	// return kv
	// return &kvstore{}
}

// type kvstore struct {
// 	//dir     string
// 	//memSize int
// }
// func (kv kvstore) Open() {
// }
// func (kv kvstore) Close() {
// }
// func (kv kvstore) Delete() {
// }
// func (kv kvstore) Put(key uint64, value [10]byte) {
// 	// kv.data[key] = value
// }
// func (kv kvstore) Get(key uint64) (value [10]byte) {
// 	return [10]byte{}
// }
