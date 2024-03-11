package kvstore

type KVStore interface {
	New() (kv KVStore, er error)
	Open()
	Delete()
	Close()
	Put(key uint64, value [10]byte)
	Get(key uint64) (value [10]byte)
}

type kvstore struct {
}

//	func (kv kvstore) New() (kv *KVStore, er error) {
//		return
//	}
func (kv kvstore) Open() {

}
func (kv kvstore) Delete() {

}
func (kv kvstore) Close() {

}
func (kv kvstore) Put(key uint64, value [10]byte) {

}
func (kv kvstore) Get(key uint64) (value [10]byte) {
	return
}

func New() (kv KVStore) {
	return
}
