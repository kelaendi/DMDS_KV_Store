package kvstore

type SSTable interface {
	Put(key uint64, value [10]byte)
	Get(key uint64) (value [10]byte)
}
