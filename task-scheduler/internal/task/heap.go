package task

type TaskHeap []*Task

// sort methods
func (h TaskHeap) Len() int { return len(h) }
func (h TaskHeap) Less(i, j int) bool {
	if h[i].ExecutionTime.Compare(h[j].ExecutionTime) == 0 {
		return h[i].Priority < h[j].Priority
	}
	return h[i].ExecutionTime.Before(h[j].ExecutionTime)
}
func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// heap methods
func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(*Task))
}
func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return item
}

func NewTaskHeap() *TaskHeap {
	return &TaskHeap{}
}
