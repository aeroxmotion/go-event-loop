package main

import (
	"fmt"
	"math/rand"

	eventLoop "github.com/aeroxmotion/go-event-loop/event_loop"
	"github.com/aeroxmotion/go-event-loop/promise"
)

func main() {
	loop := eventLoop.DefaultLoop

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

		prom := promise.New(func(resolve promise.Resolve, reject promise.Reject) {
			fmt.Println("Executing promise")

			if rand.Intn(10) > 5 {
				resolve("resolved")
			} else {
				reject("rejected")
			}
		})

		prom.
			Then(func(value any) {
				fmt.Printf("Resolved with value: %s\n", value.(string))
			}).
			Catch(func(reason any) {
				fmt.Printf("Rejected with reason: %s\n", reason.(string))
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
