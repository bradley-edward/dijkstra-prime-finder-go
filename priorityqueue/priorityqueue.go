package priorityqueue

type priorityQueueItem struct {
	Value    uint
	Priority uint
}

type PriorityQueue struct {
	items []priorityQueueItem
}

func New() PriorityQueue {
	return PriorityQueue{
		items: []priorityQueueItem{},
	}
}

func (pq *PriorityQueue) sinkDown() {
	idx := 0
	length := len(pq.items)
	element := pq.items[0]
	for {
		leftChildIdx := (2 * idx) + 1
		rightChildIdx := leftChildIdx + 1
		var swapIdx int = -1
		var leftChild, rightChild priorityQueueItem

		if leftChildIdx < length {
			leftChild = pq.items[leftChildIdx]
			if leftChild.Priority < element.Priority {
				swapIdx = leftChildIdx
			}
		}

		if rightChildIdx < length {
			rightChild = pq.items[rightChildIdx]
			if swapIdx < 0 {
				if rightChild.Priority < element.Priority {
					swapIdx = rightChildIdx
				}
			} else {
				if rightChild.Priority < leftChild.Priority {
					swapIdx = rightChildIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}

		pq.items[idx] = pq.items[swapIdx]
		pq.items[swapIdx] = element
		idx = swapIdx
	}
}

func (pq *PriorityQueue) bubbleUp() {
	idx := len(pq.items) - 1
	elem := pq.items[idx]

	for idx > 0 {
		parentIdx := (idx - 1) / 2
		parentItem := pq.items[parentIdx]

		if elem.Priority >= parentItem.Priority {
			break
		}

		pq.items[parentIdx] = elem
		pq.items[idx] = parentItem
		idx = parentIdx
	}
}

func (pq *PriorityQueue) Enqueue(val uint, prior uint) {
	pq.items = append(pq.items, priorityQueueItem{
		Value:    val,
		Priority: prior,
	})
	pq.bubbleUp()
}

func (pq *PriorityQueue) Dequeue() priorityQueueItem {
	itemsMaxIdx := len(pq.items) - 1
	minItem := pq.items[0]
	endItem := pq.items[itemsMaxIdx]
	pq.items = pq.items[:itemsMaxIdx]

	if len(pq.items) > 0 {
		pq.items[0] = endItem
		pq.sinkDown()
	}
	return minItem
}

func (pq *PriorityQueue) OverwriteMinItem(val uint, prior uint) {
	pq.items[0] = priorityQueueItem{
		Value:    val,
		Priority: prior,
	}
	pq.sinkDown()
}

func (pq PriorityQueue) Peek() priorityQueueItem {
	return pq.items[0]
}
