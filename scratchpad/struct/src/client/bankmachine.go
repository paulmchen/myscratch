package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/yireyun/go-queue"
)

// Client is our mock client struct
type Client struct {
	ID int
}

// NewClient returns a Client pointer
func NewClient(id int) *Client {
	return &Client{ID: id}
}

// DoSomething is a mock function
func (c *Client) DoSomething() {
	fmt.Printf("Client %d is operating the machine.\n", c.ID)
}

// Member is an interface describe client
type Member interface {
	DoSomething()
}

// Line is the line struct
type Line struct {
	wg   *sync.WaitGroup
	line *queue.EsQueue
	mu   *sync.Mutex
}

// NewLine returns a Line pointer
func NewLine(cap uint32) *Line {
	return &Line{
		line: queue.NewQueue(cap),
		wg:   &sync.WaitGroup{},
		mu:   &sync.Mutex{},
	}
}

// Size returns size of current line
func (l *Line) Size() uint32 {
	return l.line.Quantity()
}

// Join a member into line
func (l *Line) Join(m Member) {
	// add to wait group
	l.wg.Add(1)

	// lock the queue
	l.mu.Lock()
	defer l.mu.Unlock()

	l.line.Put(m)
}

// DoJob pop from the queue and do jobs
func (l *Line) DoJob() {
	for {
		if m, ok, newSize := l.line.Get(); ok {
			member := m.(Member)
			member.DoSomething()

			l.wg.Done()
			fmt.Printf("line size now is %d \n", newSize)
		}

		time.Sleep(time.Second * 3)
	}
}

// Wait for all goroutines done
func (l *Line) Wait() {
	l.wg.Wait()
}

func main() {
	// create a new line with capbility 5, the line goroutine will keep running, we can join line when ever we need
	l := NewLine(5)

	for i := 0; i < 10; i++ {
		c := NewClient(i)
		// put the new client into the line
		go l.Join(c)
		// fmt.Printf("add new person to the lane: %s", c.ID)
	}

	// start the line
	go l.DoJob()

	l.Wait()
}
