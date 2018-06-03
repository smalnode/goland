package lang

import (
	"errors"
	"sync"
	"testing"
	"time"
)

type Mutex interface {
	TryLock() error
	Unlock()
	IsLocked() bool
}

type mu struct {
	ch       chan struct{}
	isLocked bool
}

var ErrCannotLock = errors.New("cannot lock ")

func (mu *mu) TryLock() error {
	select {
	case mu.ch <- struct{}{}:
		mu.isLocked = true
		return nil
	default:
		return ErrCannotLock
	}
}

func (mu *mu) Unlock() {
	if mu.isLocked {
		<-mu.ch
		mu.isLocked = false
	}
}

func (mu *mu) IsLocked() bool {
	return mu.isLocked
}

func NewMutex() *mu {
	return &mu{
		ch:       make(chan struct{}, 1),
		isLocked: false,
	}
}

func TestTryLock(t *testing.T) {
	mu := NewMutex()
	v := 0
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(v *int, mu Mutex, wg *sync.WaitGroup) {
		k := 0
		for {
			<-time.After(10 * time.Millisecond)
			err := mu.TryLock()
			if err == nil {
				t.Log("#10 locked ")
				*v += 10
				<-time.After(5 * time.Millisecond)
				k++
				mu.Unlock()
				t.Log("#10 unlocked ")
			}
			if k == 10 {
				break
			}
		}
		wg.Done()
	}(&v, mu, &wg)

	go func(v *int, mu Mutex, wg *sync.WaitGroup) {
		k := 0
		for {
			<-time.After(5 * time.Millisecond)
			err := mu.TryLock()
			if err == nil {
				t.Log("#05 locked ")
				*v += 5
				<-time.After(3 * time.Millisecond)
				k++
				mu.Unlock()
				t.Log("#05 unlocked ")
			}
			if k == 10 {
				break
			}
		}
		wg.Done()
	}(&v, mu, &wg)

	wg.Wait()

	if v != 150 {
		t.Errorf("expected %d, got %d ", 150, v)
	}
}
