package mutex

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app/modules/commonm/services"
	"github.com/goatcms/goatcore/workers"
)

func TestMutexStory(t *testing.T) {
	t.Parallel()
	var waitGroup = &sync.WaitGroup{}
	for ij := workers.AsyncTestReapeat; ij > 0; ij-- {
		var (
			sharedMutex = NewSharedMutex()
		)
		waitGroup.Add(3)
		go (func() {
			unlockHandler := sharedMutex.Lock(services.LockMap{
				"a": services.LockRW,
				"b": services.LockR,
				"c": services.LockRW,
			})
			sleep()
			unlockHandler.Unlock()
			waitGroup.Done()
		})()
		go (func() {
			unlockHandler := sharedMutex.Lock(services.LockMap{
				"b": services.LockR,
				"c": services.LockRW,
				"a": services.LockRW,
			})
			sleep()
			unlockHandler.Unlock()
			waitGroup.Done()
		})()
		go (func() {
			unlockHandler := sharedMutex.Lock(services.LockMap{
				"d": services.LockR,
				"c": services.LockRW,
				"b": services.LockRW,
				"a": services.LockRW,
			})
			sleep()
			unlockHandler.Unlock()
			waitGroup.Done()
		})()
		waitGroup.Wait()
	}
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n + 10))
}
