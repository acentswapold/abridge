package block

import (
	"fmt"
	"time"

	"github.com/acentswap/abridge/log"
	"github.com/acentswap/abridge/tokens/tools"
)

var (
	scannedTxs = tools.NewCachedScannedTxs(3000)
)

// StartPoolTransactionScanJob scan job
func (b *Bridge) StartPoolTransactionScanJob() {
	chainName := b.ChainConfig.BlockChain
	log.Infof("[scanpool] start scan %v tx pool job", chainName)
	errorSubject := fmt.Sprintf("[scanpool] get %v pool txs error", chainName)
	scanSubject := fmt.Sprintf("[scanpool] scanned %v tx", chainName)
	for {
		txids, err := b.GetPoolTxidList()
		if err != nil {
			log.Error(errorSubject, "err", err)
			time.Sleep(retryIntervalInScanJob)
			continue
		}
		for _, txid := range txids {
			if scannedTxs.IsTxScanned(txid) {
				continue
			}
			log.Trace(scanSubject, "txid", txid)
			b.processTransaction(txid)
			scannedTxs.CacheScannedTx(txid)
		}
		time.Sleep(restIntervalInScanJob)
	}
}
