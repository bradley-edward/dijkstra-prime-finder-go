package main

type pqItem struct {
	value    uint
	priority uint
}

type priorityQueue struct {
	items []pqItem
}

func newPriorityQueue() priorityQueue {
	return priorityQueue{
		items: []pqItem{},
	}
}

func (pq *priorityQueue) sinkDown() {
	var idx, length int = 0, len(pq.items)

	element := pq.items[0]
	for {
		leftChildIdx := (2 * idx) + 1
		rightChildIdx := leftChildIdx + 1
		var swapIdx int = -1
		var leftChild, rightChild pqItem

		if leftChildIdx < length {
			leftChild = pq.items[leftChildIdx]
			if leftChild.priority < element.priority {
				swapIdx = leftChildIdx
			}
		}

		if rightChildIdx < length {
			rightChild = pq.items[rightChildIdx]
			if swapIdx < 0 {
				if rightChild.priority < element.priority {
					swapIdx = rightChildIdx
				}
			} else {
				if rightChild.priority < leftChild.priority {
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

func (pq *priorityQueue) bubbleUp() {
	idx := len(pq.items) - 1
	elem := pq.items[idx]

	for idx > 0 {
		parentIdx := (idx - 1) / 2
		parentItem := pq.items[parentIdx]

		if elem.priority >= parentItem.priority {
			break
		}

		pq.items[parentIdx] = elem
		pq.items[idx] = parentItem
		idx = parentIdx
	}
}

func (pq *priorityQueue) enqueue(val, prior uint) {
	pq.items = append(pq.items, pqItem{
		value:    val,
		priority: prior,
	})
	pq.bubbleUp()
}

/*
func (pq *priorityQueue) dequeue() pqItem {
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
*/

func (pq *priorityQueue) overwriteMinItem(val, prior uint) {
	pq.items[0] = pqItem{
		value:    val,
		priority: prior,
	}
	pq.sinkDown()
}

func (pq *priorityQueue) peek() pqItem {
	return pq.items[0]
}

func dijkstraAlgo(inputValue uint) *[]uint {
	primesList := []uint{2}

	multiplesPq := newPriorityQueue()
	multiplesPq.enqueue(2, 4)

	var num uint
	for num = 3; num <= inputValue; num++ {
		// Find the smallest "multiple" in the PQ.
		minMultiple := multiplesPq.peek()
		// If the current number is less than the smallest multiple, the current number is a prime.
		if num < minMultiple.priority {
			primesList = append(primesList, num)
			multiplesPq.enqueue(num, num*num)
			continue
		}

		// Else, the number is NOT a prime.
		for num == minMultiple.priority {
			multiplesPq.overwriteMinItem(minMultiple.value, minMultiple.value+minMultiple.priority)
			minMultiple = multiplesPq.peek()
		}
	}

	return &primesList
}