package context_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"
)

const MAX_GOROUTIN = 4
const DIFFICULTY = 5

func TestMining(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int)
	for i := 0; i < MAX_GOROUTIN; i++ {
		go func(ctx context.Context, from, step, difficulty int, ch chan<- int) {
			for n := from; true; n += step {
				select {
				case <-ctx.Done():
					return
				default:
					if check(n, difficulty) {
						select {
						case ch <- n:
							close(ch)
						case <-ctx.Done():
						}
						return
					}
				}
			}
		}(ctx, i, MAX_GOROUTIN, DIFFICULTY, ch)
	}

	select {
	case nounce := <-ch:
		cancel()
		b := sha256.Sum256([]byte(string(nounce)))
		sum := hex.EncodeToString(b[:])
		t.Logf("sha256.sum(%d) = %s\n", nounce, sum)
	case <-time.After(10 * time.Second):
		t.Log("timeout!!!")
		cancel()
	}
}

func check(nounce, difficulty int) bool {
	b := sha256.Sum256([]byte(string(nounce)))
	sum := hex.EncodeToString(b[:])
	for i := 0; i < difficulty; i++ {
		if sum[i] != '0' {
			return false
		}
	}
	return true
}
