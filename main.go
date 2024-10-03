package main

import (
	"fmt"

	eventLoop "github.com/aeroxmotion/go-event-loop/event_loop"
)

func main() {
	loop := eventLoop.New()

	loop.Exec(func(_ ...any) {
		fmt.Println("1")

		loop.Exec(func(_ ...any) {
			fmt.Println("Immediate 1")

			loop.Exec(func(_ ...any) {
				fmt.Println("Immediate 2")
			})
		})

		loop.Exec(func(_ ...any) {
			fmt.Println("Immediate 3")
		})

		loop.QueueTask(func(_ ...any) {
			fmt.Println("5")
		})

		loop.QueueMicroTask(func(_ ...any) {
			fmt.Println("3")
		})

		loop.QueueMicroTask(func(_ ...any) {
			fmt.Println("4")
		})

		fmt.Println("2")
	})

	fmt.Println("Hello world!")
}
