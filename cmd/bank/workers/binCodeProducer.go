package workers

import "sync"

// BinCodeProducer 生成银行的bin号码
// 100000 - 999999 涵盖6位所有数字 89W个

type BinCodeProducer struct {
	Start uint64
	End   uint64
}

// GenerateBinCodeQueue 产生需要消化的银行bin码
// wg 控制线程
func (b *BinCodeProducer) GenerateBinCodeQueue(wg *sync.WaitGroup, queue chan<- uint64) {
	for i := b.Start; i <= b.End; i++ {
		wg.Add(1)
		queue <- i
	}
}
