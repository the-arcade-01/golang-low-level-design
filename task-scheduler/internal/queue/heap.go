package queue

import "task-scheduler/internal/models"

type Heap []*models.Task

// sort interface
func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	if h[i].StartTime.Compare(h[j].StartTime) == 0 {
		return h[i].Priority < h[j].Priority // High = 0, Medium = 1, Low = 2
	}
	return h[i].StartTime.Before(h[j].StartTime)
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// heap methods
func (h *Heap) Push(x any) {
	*h = append(*h, x.(*models.Task))
}
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return item
}

func NewHeap() *Heap {
	return &Heap{}
}
