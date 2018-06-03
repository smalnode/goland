package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGQUIT)

	cond := &sync.Cond{L: &sync.Mutex{}}

	go subscribe(cond, "one")
	go subscribe(cond, "two")
	go subscribe(cond, "three")

	for {
		select {
		case <-ch:
			os.Exit(0)
		default:
			input(cond)
		}
	}
}

func input(c *sync.Cond) {
	fmt.Scanf("%n")
	c.Broadcast()
}

func subscribe(c *sync.Cond, m string) {
	for {
		c.L.Lock()
		c.Wait()
		fmt.Println(m)
		c.L.Unlock()
	}
}
