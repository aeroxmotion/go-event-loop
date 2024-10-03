package callStack

import "github.com/aeroxmotion/go-event-loop/task"

type CallStack []*task.Task

func New() *CallStack {
	return &CallStack{}
}

func (stack *CallStack) HasTasks() bool {
	return len(*stack) > 0
}

func (stack *CallStack) Push(task *task.Task) {
	*stack = append(*stack, task)
}

func (stack *CallStack) Pop() {
	*stack = (*stack)[:len(*stack)-1]
}
