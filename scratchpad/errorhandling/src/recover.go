package main

import (
	"errors"
	"fmt"
)

func yikes() {
	panic(errors.New("Something bad happened"))
}

func main() {
	msg := "Everything is fine"
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
		fmt.Println(msg)
	}()

	yikes()
}
