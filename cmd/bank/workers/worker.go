package workers

import (
	"log"
	"sync"
)

var workerCount = 20

// Work 开启线程 工作
func Work(queue <-chan uint64, group *sync.WaitGroup, m *ProxyMan) {
	if len(m.AddressList) < workerCount {
		log.Panicf("代理地址数量比worker还少")
		return
	}
	for i := 0; i < workerCount; i++ {
		go func() {
			DoSomething(queue, group, m.AddressList[i])
		}()
	}
}

// DoSomething 工作
func DoSomething(queue <-chan uint64, group *sync.WaitGroup, address string) {
	requester := &Requester{
		ProxyAddress: address,
	}

	for q := range queue {
		requester.RequestBankInfo(q)
		group.Done()
	}
}
