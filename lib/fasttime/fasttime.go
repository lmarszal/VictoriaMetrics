package fasttime

import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"sync/atomic"
	"time"
)

func init() {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for tm := range ticker.C {
			t := uint64(tm.Unix())
			atomic.StoreUint64(&currentTimestamp, t)
			// DEBUG log mocked timestamp every 10s so that we know "when" we are
			{
				tmp := UnixTimestamp()
				if tmp%10 == 0 {
					logger.Infof("Now is %d", tmp)
				}
			}
		}
	}()
}

var currentTimestamp = uint64(time.Now().Unix())

// DEBUG capture app start time for time traveling
var debugStartTimestamp = uint64(time.Now().Unix())

// UnixTimestamp returns the current unix timestamp in seconds.
//
// It is faster than time.Now().Unix()
func UnixTimestamp() uint64 {
	// DEBUG time travel - simulate that the app started at 2022-10-27 23:59:45 UTC (1666915185)
	return atomic.LoadUint64(&currentTimestamp) - debugStartTimestamp + 1666915185
	//return atomic.LoadUint64(&currentTimestamp)
}

// UnixDate returns date from the current unix timestamp.
//
// The date is calculated by dividing unix timestamp by (24*3600)
func UnixDate() uint64 {
	return UnixTimestamp() / (24 * 3600)
}

// UnixHour returns hour from the current unix timestamp.
//
// The hour is calculated by dividing unix timestamp by 3600
func UnixHour() uint64 {
	return UnixTimestamp() / 3600
}
