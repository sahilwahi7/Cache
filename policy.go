package main

import (
	"Cache/Storagefactory"
	"fmt"
)

func main() {
	storage := &Storagefactory.HashmapStorage{
		Hashstore: map[int]int{},
		Capacity:  10,
	}
	storage.Put(1, 4)
	storage.Put(3, 5)
	storage.Put(2, 4)
	fmt.Println(storage.Hashstore)
	storage.Get(3)
	storage.Get(5)

}
