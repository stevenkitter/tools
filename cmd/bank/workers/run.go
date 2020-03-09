package workers

import (
	"log"
	"sync"
)

func Run() {
	p := NewProxyMan()
	for {
		p.RequestAddressList()
		if len(p.AddressList) > 20 {
			log.Printf("代理ip地址已有20个 可以使用")
			break
		}
	}

	wg := &sync.WaitGroup{}
	queue := make(chan uint64, 0)
	b := BinCodeProducer{Start: 100000, End: 999999}
	go b.GenerateBinCodeQueue(wg, queue)
	Work(queue, wg, &p)
	//
	wg.Wait()

}
