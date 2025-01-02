package utils

import "math"

func BlockToEpoch(blockOffset, epochSize int64) int64 {
	epoch := math.Ceil(float64(blockOffset) / float64(epochSize))
	return int64(epoch)
}
