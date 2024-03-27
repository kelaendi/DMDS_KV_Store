package kvstore

import (
	"math/rand"
)

type Skiplist interface {
	Put(key uint64, value [10]byte)
	Get(key uint64) (value [10]byte)
}

func NewSkiplist(maxHeight int) (sl Skiplist) {
	sl = &skiplist{
		maxHeight: maxHeight,
		head: head{
			next:   nil,
			height: 0,
		},
	}
	return sl
}

type node struct {
	key    uint64
	value  [10]byte
	next   *node
	prev   *node
	height int
}

type head struct {
	next   *node
	height int
}

type skiplist struct {
	maxHeight int
	head      head
}

func (sl skiplist) Put(key uint64, value [10]byte) {
	node := node{
		key:    key,
		value:  value,
		next:   nil,
		prev:   nil,
		height: 0,
	}
	sl.insertNode(&node)

	node.height = randomHeight(sl.maxHeight)
}
func (sl skiplist) Get(key uint64) (value [10]byte) {
	//TODO:
	return value
}

func randomHeight(maxHeight int) (height int) {
	done := false
	for height < maxHeight && !done {
		if rand.Intn(2) == 0 {
			done = true
		} else {
			height += 1
		}
	}
	return height
}

func (sl skiplist) insertNode(node *node) {
	currentNode := sl.head.next
	currentHeight := sl.head.height

	for currentNode != nil && node.key > currentNode.key {
		if currentNode.next == nil || currentNode.key >= node.key {
			if currentHeight == 0 {
				currentNode.next = node
				node.prev = currentNode
				break
			} else {
				currentHeight -= 1
			}
		} else {
			currentNode = currentNode.next
		}
	}
}
