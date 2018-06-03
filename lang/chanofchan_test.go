package lang

import "testing"
import "time"

type bus chan chan int

func TestChanofChan(t *testing.T) {
	b5 := make(bus, 5)

	go func() {
		for i := 0; i < 6; i++ {
			c := make(chan int)
			b5 <- c
			go func(v int) {
				i++ // i is a reference
				for j := 0; j < 10; j++ {
					c <- v
				}
				close(c)
			}(i)
		}
		close(b5)
	}()

	for c := range b5 {
		t.Logf("receive a chan from bus %v ", c)
		for v := range c {
			t.Logf("receive a int from chan %d ", v)
		}
	}
}

func TestChan(t *testing.T) {
	c := make(chan int)
	c2 := make(chan int)
	b := make(bus)

	go func() {
		b <- c
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		b <- c2
		for i := 0; i < 10; i++ {
			c2 <- i + 10
		}
		close(c2)
	}()

	go func() {
		for {
			c, ok := <-b
			if !ok {
				break
			}
			for v := range c {
				t.Log(v)
			}
		}
	}()

	<-time.After(5 * time.Second)
}

func TestGoRef(t *testing.T) {
	v := 0
	c := make(chan struct{})
	go func() {
		v += 2
		c <- struct{}{}
	}()
	<-c
	if v == 2 {
		t.Log("v in goroutine is a reference ")
	} else {
		t.Log("v in goroutine is a value copy ")
	}
}
