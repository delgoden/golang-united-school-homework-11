package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	sem := make(chan struct{}, pool)
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for i := int64(0); i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(locker *sync.Mutex, wg *sync.WaitGroup, id int64) {
			user := getOne(id)
			locker.Lock()
			res = append(res, user)
			locker.Unlock()
			<-sem
			wg.Done()
		}(mu, wg, i)
	}
	wg.Wait()
	return res
}
