package algorithm

type Lru struct {
	lruMap map[int]*Node
	head   *Node
	tail   *Node
	size   int
}

type Node struct {
	key   int
	value string
	pre   *Node
	next  *Node
}

func NewLru(size int) *Lru {
	return &Lru{
		lruMap: make(map[int]*Node),
		size:   size,
	}
}

func (lru *Lru) Put(key int, value string) *Node {
	node := &Node{
		key:   key,
		value: value,
	}
	lru.moveToHead(node)
	lru.lruMap[key] = node
	if len(lru.lruMap) > lru.size {
		delete(lru.lruMap, lru.tail.key)
		lru.tail.pre.next = nil
		lru.tail = lru.tail.pre
	}
	return node
}

func (lru *Lru) Get(key int) *Node {
	node := lru.lruMap[key]
	if node == nil {
		return nil
	}
	lru.moveToHead(node)
	return node
}

func (lru *Lru) moveToHead(node *Node) {
	// 如果链表为空，则将head和tail初始化即可
	if lru.head == nil {
		lru.head = node
		lru.tail = node
		return
	}
	// 如果node是头节点，无需任何操作
	if lru.head == node {
		return
	}
	// 如果node是尾节点，将node节点的前节点设为尾节点
	if lru.tail == node {
		lru.tail.pre.next = nil
		lru.tail = lru.tail.pre
	}
	// 将node节点设置为头节点
	lru.head.pre = node
	node.next = lru.head
	lru.head = node
}
