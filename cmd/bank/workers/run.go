package workers

import (
	"log"
	"sync"
)

// Run run
func Run() {
	p := NewProxyMan()

	p.RequestAddressList()
	if len(p.AddressList) > 5 {
		log.Printf("代理ip地址已有20个 可以使用")
	}

	wg := &sync.WaitGroup{}
	queue := make(chan uint64, 0)
	b := BinCodeProducer{Start: 100000, End: 999999}
	Work(queue, wg, &p)
	b.GenerateBinCodeQueue(wg, queue)

	//
	wg.Wait()
	log.Printf("打完收工了，厉害啊！")
}
