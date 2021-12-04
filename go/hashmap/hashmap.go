package hashmap

type pair struct {
	key   string
	value int
}

type node struct {
	data pair
	next *node
}

type linkedList struct {
	head *node
}

const arraySize = 100

type hashmap struct {
	array [arraySize]*linkedList
}

func NewHashmap() *hashmap {
	var result hashmap
	for i := range result.array {
		result.array[i] = NewLinkedList()
	}
	return &result
}

func (h *hashmap) Put(key string, value int) {
	newPair := pair{key, value}
	hashCode := hash(key)
	h.array[hashCode].Prepend(newPair)
}

func (h *hashmap) Get(key string) (int, bool) {
	hashCode := hash(key)
	return h.array[hashCode].Search(key)
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) contains(pairKey string) bool {
	if l.head == nil {
		return false
	}

	currentNode := l.head
	for currentNode != nil {
		if currentNode.data.key == pairKey {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (l *linkedList) Prepend(data pair) {
	if l.contains(data.key) {
		return
	}
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *linkedList) Search(pairKey string) (int, bool) {
	if l.head == nil {
		return 0, false
	}

	currentNode := l.head
	for currentNode != nil {
		if currentNode.data.key == pairKey {
			return currentNode.data.value, true
		}
		currentNode = currentNode.next
	}
	return 0, false
}

func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}
