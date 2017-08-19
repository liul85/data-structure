package queue

type node struct {
	data interface{}
	next *node
}

type queue struct {
	head *node
	tail *node
}

func New(args ...interface{}) *queue {
	queue := &queue{}
	for _, v := range args {
		queue.Enqueue(v)
	}

	return queue
}

func (queue *queue) Enqueue(data interface{}) {
	node := &node{data: data}
	if queue.head == nil && queue.tail == nil {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
}
