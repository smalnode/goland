package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	for {
		select {
		case <-ch:
			fmt.Println("received syscall.SIGKILL ")
			os.Exit(0)
		default:
			scanMultiply()
		}
	}
}

func scanMultiply() {
	fmt.Println("Input two integers(CTRL-C to exit): ")
	var lhs, rhs int64
	_, err := fmt.Scanf("%d%d\n", &lhs, &rhs)
	if err == nil {
		fmt.Println(lhs * rhs)
	}
}
