package Storagefactory

import (
	"fmt"
)

type HashmapStorage struct {
	Storage
	Hashstore map[int]int //using sync map to prevent race conditions
	Capacity  int
}

func (h *HashmapStorage) isStorageFull() bool {
	if len(h.Hashstore) == h.Capacity {
		return true
	}
	return false
}
func (h *HashmapStorage) Get(key int) int {
	val := h.Hashstore[key]
	return val

}

func (h *HashmapStorage) Put(key int, val int) {
	if h.isStorageFull() {
		fmt.Println("Storage full")
		return
	}
	h.Hashstore[key] = val
}
