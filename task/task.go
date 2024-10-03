package task

type Runner = func(...any)

type Task struct {
	Args    []any
	Runner  Runner
	Running bool
}

func New(runner Runner, args []any) *Task {
	return &Task{
		Args:    args,
		Runner:  runner,
		Running: false,
	}
}

func (task *Task) Run() {
	task.Running = true
	task.Runner(task.Args...)
	task.Running = false
}
