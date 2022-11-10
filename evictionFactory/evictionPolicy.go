package evictionFactory

// interface of eviction policy, which can be used by multiple polices
type evictionPolicy interface {
	keyAccessed(key int)
	evictKey()
}
