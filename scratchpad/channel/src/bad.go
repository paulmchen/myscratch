package main

import (
	"fmt"
	"time"
)

func send(ch chan string) {

	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	go send(msg)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			close(msg)
			time.Sleep(500 * time.Millisecond)
			return
		}
	}

}
