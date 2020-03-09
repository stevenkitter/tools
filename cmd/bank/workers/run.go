package workers

import (
	"log"
	"sync"
)

func Run() {
	p := NewProxyMan()
	for {
		p.RequestAddressList()
		if len(p.AddressList) > 5 {
			log.Printf("代理ip地址已有20个 可以使用")
			break
		}
	}

	wg := &sync.WaitGroup{}
	queue := make(chan uint64, 0)
	b := BinCodeProducer{Start: 621466, End: 999999}
	Work(queue, wg, &p)
	b.GenerateBinCodeQueue(wg, queue)

	//
	wg.Wait()

}
