package lang

import (
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(v int, wg sync.WaitGroup) {
			t.Log(v)
			wg.Done()
		}(i, wg)
	}

	wg.Wait()

}

func TestMutexPanic(t *testing.T) {
	mu := sync.Mutex{}

	go func() {
		mu.Lock()
		defer mu.Unlock()

		t.Log("lock 2")
		<-time.After(1 * time.Second)
		panic("hahahahah  ")
	}()

	go func() {
		<-time.After(100 * time.Millisecond)
		mu.Lock()
		defer mu.Lock()

		t.Log("lock 2")
	}()

	<-time.After(2 * time.Second)
}
