package models

// QueueNode -
type QueueNode struct {
	priority float64
	data     interface{}
}

// PriorityQueue -
type PriorityQueue struct {
	arr         []QueueNode
	Size        int
	GetPriority func(interface{}) float64
}

// Enqueue -
func (m *PriorityQueue) Enqueue(data interface{}) {
	priority := m.GetPriority(data)
	node := QueueNode{data: data, priority: priority}
	m.arr = append(m.arr, node)
	m.Size = len(m.arr)
	m.bubbleUp()
}

// Dequeue -
func (m *PriorityQueue) Dequeue() interface{} {
	if len(m.arr) == 0 {
		return nil
	}
	node := m.arr[0]
	m.swap(0, m.Size-1)
	// remove last element
	m.arr = m.arr[:len(m.arr)-1]
	m.Size = len(m.arr)
	m.bubbleDown()
	return node.data
}

func (m *PriorityQueue) bubbleUp() {
	i := m.Size - 1
	parent := getParent(i)
	for parent >= 0 && (m.arr[i].priority > m.arr[parent].priority) {
		m.swap(i, parent)
		i = parent
		parent = getParent(parent)
	}
}

func (m *PriorityQueue) bubbleDown() {
	i := 0
	child := m.getLargestChild(i)
	for child > 0 && (m.arr[i].priority < m.arr[child].priority) {
		m.swap(i, child)
		i = child
		child = m.getLargestChild(child)
	}
}

func (m *PriorityQueue) swap(a, b int) {
	var tmp QueueNode
	tmp = m.arr[a]
	m.arr[a] = m.arr[b]
	m.arr[b] = tmp
}

func (m *PriorityQueue) getLargestChild(i int) int {
	l := getLeftChild(i)
	r := getRightChild(i)
	if l >= m.Size && r >= m.Size {
		return -1
	}
	if l >= m.Size {
		return r
	}
	if r >= m.Size {
		return l
	}

	if m.arr[l].priority < m.arr[r].priority {
		return r
	}
	return l
}

func getParent(i int) int {
	return i / 2
}

func getLeftChild(i int) int {
	return (i * 2)
}

func getRightChild(i int) int {
	return (i * 2) + 1
}
