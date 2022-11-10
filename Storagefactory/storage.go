package Storagefactory

type Storage interface {
	get(int) int
	put(int, int)
}
