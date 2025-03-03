package priorityqueue

type priorityQueueItem struct {
	value    int
	priority int
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

func (pq *PriorityQueue) bubbleUp() {
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

func (pq *PriorityQueue) Enqueue(val, prior int) {
	pq.items = append(pq.items, priorityQueueItem{
		value:    val,
		priority: prior,
	})
	pq.bubbleUp()
}

func (pq *PriorityQueue) Dequeue() priorityQueueItem {
	/*
	   min_item = self.items[0]
	   end_item = self.items.pop()
	   if len(self.items) > 0:
	       self.items[0] = end_item
	       self.sink_down()
	   return min_item
	*/

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

func (pq PriorityQueue) Peek() priorityQueueItem {
	return pq.items[0]
}

/*
   def __init__(self):
       self.items = []

   def sink_down(self):
       idx = 0
       length = len(self.items)
       element = self.items[0]
       while True:
           left_child_idx = (2 * idx) + 1
           right_child_idx = left_child_idx + 1
           left_child = None
           right_child = None
           swap = None

           if left_child_idx < length:
               left_child = self.items[left_child_idx]
               if left_child[1] < element[1]:
                   swap = left_child_idx

           if right_child_idx < length:
               right_child = self.items[right_child_idx]
               if (swap is None and right_child[1] < element[1]) or (swap is not None and right_child[1] < left_child[1]):
                  swap = right_child_idx

           if swap is None:
               break

           self.items[idx] = self.items[swap]
           self.items[swap] = element
           idx = swap

   def bubble_up(self):
       idx = len(self.items) - 1
       element = self.items[idx]

       while idx > 0:
           parent_idx = (idx - 1) // 2
           parent_item = self.items[parent_idx]
           if element[1] >= parent_item[1]:
               break

           self.items[parent_idx] = element
           self.items[idx] = parent_item
           idx = parent_idx

   def enqueue(self, value, priority):
       self.items.append(
           (value, priority)
       )
       self.bubble_up()

   def dequeue(self):
       min_item = self.items[0]
       end_item = self.items.pop()
       if len(self.items) > 0:
           self.items[0] = end_item
           self.sink_down()
       return min_item

   def overwrite_min_item(self, value, priority):
       self.items[0] = (value, priority)
       self.sink_down()

   def peek(self):
       return self.items[0]
*/