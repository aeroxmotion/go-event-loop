package promise

import eventLoop "github.com/aeroxmotion/go-event-loop/event_loop"

type Resolve = func(value any)
type Reject = func(reason any)

type Executor = func(
	resolve Resolve,
	reject Reject,
)

type Promise struct {
	then  Resolve
	catch Reject
}

func New(executor Executor) *Promise {
	loop := eventLoop.DefaultLoop
	promise := &Promise{}

	resolve := func(value any) {
		loop.QueueMicroTask(func(a ...any) {
			promise.then(value)
		})
	}

	reject := func(reason any) {
		loop.QueueMicroTask(func(a ...any) {
			promise.catch(reason)
		})
	}

	loop.Exec(func(a ...any) {
		executor(resolve, reject)
	})

	return promise
}

func (promise *Promise) Then(then Resolve) *Promise {
	promise.then = then
	return promise
}

func (promise *Promise) Catch(catch Reject) {
	promise.catch = catch
}
