package main

import (
	"errors"
	"time"

	"./safely"
)

func message() {
	println("Inside goroutine")
	panic(errors.New("Oops!"))
}

func main() {
	safely.Go(message)
	println("Outside gorountine")
	time.Sleep(200 * time.Millisecond)
}
