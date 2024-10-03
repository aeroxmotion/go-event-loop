package taskQueue

import "github.com/aeroxmotion/go-event-loop/task"

type TaskQueue []*task.Task

func New() *TaskQueue {
	return &TaskQueue{}
}

func (queue *TaskQueue) HasTasks() bool {
	return len(*queue) > 0
}

func (queue *TaskQueue) Push(task *task.Task) {
	*queue = append(*queue, task)
}

func (queue *TaskQueue) Dequeue() *task.Task {
	resultTask := (*queue)[0]
	*queue = (*queue)[1:]

	return resultTask
}
