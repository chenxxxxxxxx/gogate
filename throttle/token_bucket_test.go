package throttle

import (
	"fmt"
	"testing"
	"time"

)

func TestNewRateLimiter(t *testing.T) {
	rl := NewRateLimiter(1000)
	fmt.Println(rl)

}

func TestRateLimiter_Acquire(t *testing.T) {
	rl := NewRateLimiter(1)

	count := 0
	for {
		rl.Acquire()
		fmt.Println(time.Now())

		count++
		if count >= 10 {
			break
		}
	}
}

func TestRateLimiter_TryAcquire(t *testing.T) {
	count := 0

	rl := NewRateLimiter(1)
	for {
		pass := rl.TryAcquire()
		if pass {
			fmt.Println(time.Now())
		}

		count++
		if count >= 10 {
			break
		}
	}

}
