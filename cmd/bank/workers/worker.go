package workers

import (
	"math/rand"
	"sync"
)

var workerCount = 20

// Work 开启线程 工作
func Work(queue <-chan uint64, group *sync.WaitGroup, m *ProxyMan) {
	for i := 0; i < workerCount; i++ {
		go func() {
			DoSomething(queue, group, m)
		}()
	}
}

func DoSomething(queue <-chan uint64, group *sync.WaitGroup, m *ProxyMan) {
	for q := range queue {
		requester := Requester{
			ProxyAddress: m.AddressList[rand.Intn(workerCount)],
		}
		requester.RequestBankInfo(q)
		group.Done()
	}
}
