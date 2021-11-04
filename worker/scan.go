package worker

import (
	"github.com/acentswap/abridge/tokens"
	"github.com/acentswap/abridge/tokens/btc"
)

// StartScanJob scan job
func StartScanJob(isServer bool) {
	srcChainCfg := tokens.SrcBridge.GetChainConfig()
	if srcChainCfg.EnableScan && btc.BridgeInstance != nil {
		go btc.BridgeInstance.StartChainTransactionScanJob()
		if srcChainCfg.EnableScanPool {
			go btc.BridgeInstance.StartPoolTransactionScanJob()
		}
		go btc.BridgeInstance.StartSwapHistoryScanJob()
	}
}
