package mutex

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
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
			unlockHandler := sharedMutex.Lock(commservices.LockMap{
				"a": commservices.LockRW,
				"b": commservices.LockR,
				"c": commservices.LockRW,
			})
			sleep()
			unlockHandler.Unlock()
			waitGroup.Done()
		})()
		go (func() {
			unlockHandler := sharedMutex.Lock(commservices.LockMap{
				"b": commservices.LockR,
				"c": commservices.LockRW,
				"a": commservices.LockRW,
			})
			sleep()
			unlockHandler.Unlock()
			waitGroup.Done()
		})()
		go (func() {
			unlockHandler := sharedMutex.Lock(commservices.LockMap{
				"d": commservices.LockR,
				"c": commservices.LockRW,
				"b": commservices.LockRW,
				"a": commservices.LockRW,
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
