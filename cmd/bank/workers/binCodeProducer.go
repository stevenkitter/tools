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
	// 完成基本的结构
	//for i := b.Start; i <= b.End; i++ {
	//	wg.Add(1)
	//	queue <- i
	//}
	// 请求数据库未完成的 失败的
	row, _ := db.Raw("SELECT bin_code FROM tools.error_bank_bins").Rows()
	defer row.Close()
	errBins := make([]uint64, 0)
	for row.Next() {
		var binCode struct {
			BinCode uint64
		}
		db.ScanRows(row, &binCode)
		errBins = append(errBins, binCode.BinCode)
	}
	for _, b := range errBins {
		wg.Add(1)
		queue <- b
	}
}
