package evictionFactory

import (
	"Cache/algorithm"
)

type LRUevictionPolicy struct {
	evictionPolicy
	dll    algorithm.Dll
	mapper map[int]*algorithm.LinkedListNode
}

func (l *LRUevictionPolicy) keyAccessed(key int) {
	if val, ok := l.mapper[key]; !ok {
		l.dll.DetachNode(val)
		l.dll.AddNodetoLast(val)
	} else {
		newNode := l.dll.AddElementAtLast(key)
		l.mapper[key] = newNode
	}
}

func (l *LRUevictionPolicy) evictKey() {
	first := l.dll.GetFirstNode()

	l.dll.DetachNode(first)
}
