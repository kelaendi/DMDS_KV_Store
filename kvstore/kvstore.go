package kvstore

type KVStore interface {
	Put(key uint64, value [10]byte)
	Get(key uint64) (value [10]byte)
}

type kvstoreControl interface {
	Open(kv KVStore) error
	Close(kv KVStore) error
	Delete(kv KVStore) error
}

type Option func(kv KVStore) KVStore

func NewKVStore(options ...Option) (kv KVStore) {
	// for _, o := range options {
	// 	kv = o(kv)
	// }
	// return &kvstore{}
	return kv
}

func newKVStoreControl() (controller kvstoreControl) {
	return controller
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
