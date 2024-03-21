package kvstore

type Skiplist interface {
	Put(key uint64, value [10]byte)
	Get(key uint64) (value [10]byte)
}
