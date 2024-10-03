package eventLoop

import (
	callStack "github.com/aeroxmotion/go-event-loop/call_stack"
	"github.com/aeroxmotion/go-event-loop/task"
	taskQueue "github.com/aeroxmotion/go-event-loop/task_queue"
)

type EventLoop struct {
	callStack      *callStack.CallStack
	taskQueue      *taskQueue.TaskQueue
	microTaskQueue *taskQueue.TaskQueue
}

func New() *EventLoop {
	return &EventLoop{
		callStack:      callStack.New(),
		taskQueue:      taskQueue.New(),
		microTaskQueue: taskQueue.New(),
	}
}

func (loop *EventLoop) Exec(runner task.Runner, args ...any) {
	loop.runSync(task.New(runner, args))
}

func (loop *EventLoop) QueueTask(runner task.Runner, args ...any) {
	loop.taskQueue.Push(task.New(runner, args))
}

func (loop *EventLoop) QueueMicroTask(runner task.Runner, args ...any) {
	loop.microTaskQueue.Push(task.New(runner, args))
}

func (loop *EventLoop) runSync(taskToRun *task.Task) {
	loop.callStack.Push(taskToRun)
	taskToRun.Run()
	loop.callStack.Pop()

	if loop.callStack.HasTasks() {
		return
	}

	if loop.microTaskQueue.HasTasks() {
		loop.runSync(loop.microTaskQueue.Dequeue())
		return
	}

	if loop.taskQueue.HasTasks() {
		loop.runSync(loop.taskQueue.Dequeue())
		return
	}
}
