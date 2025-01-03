package utils

import "math"

func BlockToEpoch(blockOffset int64, epochSize int64) int64 {
	epoch := math.Ceil(float64(blockOffset) / float64(epochSize))
	return int64(epoch)
}

func EpochToBlockRange(epoch int64, epochSize int64) (startBlock int64, endBlock int64) {
	startBlock = epoch * epochSize
	endBlock = startBlock + epochSize
	return startBlock, endBlock
}
