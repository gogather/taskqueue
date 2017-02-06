package taskqueue

import (
	"container/list"

	"errors"

	"sync"

	"github.com/gogather/safemap"
)

// TaskQueue task queue
type TaskQueue struct {
	sync.RWMutex
	m *safemap.SafeMap
	l *list.List
}

// New new a TaskQueue
func New() *TaskQueue {
	return &TaskQueue{
		m: safemap.New(),
		l: list.New().Init(),
	}
}

// Top get top element of TaskQueue
func (tq *TaskQueue) Top() (bool, string, interface{}) {
	defer func() {
		tq.RUnlock()
	}()
	tq.RLock()

	e := tq.l.Front()
	if e == nil {
		return false, "", nil
	}

	taskID := e.Value.(string)
	value, _ := tq.m.Get(taskID)

	tq.l.Remove(e)
	tq.m.Remove(taskID)

	return true, taskID, value
}

// Add add element into TaskQueue
func (tq *TaskQueue) Add(taskID string, task interface{}) {
	defer func() {
		tq.Unlock()
	}()
	tq.Lock()

	tq.m.Put(taskID, task)
	tq.l.PushBack(taskID)
	return
}

// Remove remove element from TaskQueue
func (tq *TaskQueue) Remove(taskID string) error {
	defer func() {
		tq.Unlock()
	}()
	tq.Lock()

	for e := tq.l.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == taskID {
			tq.l.Remove(e)
			tq.m.Remove(taskID)
			return nil
		}
	}
	return errors.New("remove failed")
}

// Length get length of TaskQueue
func (tq *TaskQueue) Length() int64 {
	return (int64)(tq.l.Len())
}
