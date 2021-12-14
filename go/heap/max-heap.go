package heap

type MaxHeap struct {
	Elements []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(data int) {
	h.Elements = append(h.Elements, data)
	newElementIndex := len(h.Elements) - 1
	parentIndex := getParentIndex(newElementIndex)
	for parentIndex != -1 && h.Elements[newElementIndex] > h.Elements[parentIndex] {
		h.Elements[newElementIndex], h.Elements[parentIndex] = h.Elements[parentIndex], h.Elements[newElementIndex]
		newElementIndex = parentIndex
		parentIndex = getParentIndex(newElementIndex)
	}
}

func (h *MaxHeap) Extract() (int, bool) {
	if len(h.Elements) == 0 {
		return 0, false
	}
	extractedElement := h.Elements[0]
	lastElementIndex := len(h.Elements) - 1
	h.Elements[0] = h.Elements[lastElementIndex]
	h.Elements = h.Elements[:lastElementIndex]
	parentIndex := 0
	leftChildIndex := getLeftChildIndex(parentIndex)
	rightChildIndex := getRightChildIndex(parentIndex)

	for leftChildIndex < len(h.Elements) && rightChildIndex < len(h.Elements) {
		maxChildIndex := h.max(leftChildIndex, rightChildIndex)
		if h.Elements[maxChildIndex] > h.Elements[parentIndex] {
			h.Elements[maxChildIndex], h.Elements[parentIndex] = h.Elements[parentIndex], h.Elements[maxChildIndex]
			parentIndex = maxChildIndex
			leftChildIndex = getLeftChildIndex(parentIndex)
			rightChildIndex = getRightChildIndex(parentIndex)
		}
	}

	if leftChildIndex < len(h.Elements) {
		if h.Elements[leftChildIndex] > h.Elements[parentIndex] {
			h.Elements[leftChildIndex], h.Elements[parentIndex] = h.Elements[parentIndex], h.Elements[leftChildIndex]
		}
	}
	return extractedElement, true
}

func (h *MaxHeap) max(leftChildIndex, rightChildIndex int) int {
	if h.Elements[leftChildIndex] > h.Elements[rightChildIndex] {
		return leftChildIndex
	}
	return rightChildIndex
}

func getParentIndex(childIndex int) int {
	if childIndex == 0 {
		return -1
	}
	return (childIndex - 1) / 2
}

func getLeftChildIndex(parentIndex int) int {
	return parentIndex*2 + 1
}

func getRightChildIndex(parentIndex int) int {
	return parentIndex*2 + 2
}
