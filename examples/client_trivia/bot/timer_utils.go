package bot

import (
	"math/rand/v2"
	"time"
)

// closeTimer should be used as a deferred function
// in order to cleanly shut down a timer
func closeTimer(timer *time.Timer, drained *bool) {
	if drained == nil {
		panic("drained bool pointer is nil")
	}
	if !timer.Stop() {
		if *drained {
			return
		}
		<-timer.C
		*drained = true
	}
}

// resetTimer sets drained to false after resetting the timer.
func resetTimer(timer *time.Timer, duration time.Duration, drained *bool) {
	if drained == nil {
		panic("drained bool pointer is nil")
	}
	if !timer.Stop() {
		if !*drained {
			<-timer.C
		}
	}
	timer.Reset(duration)
	*drained = false
}

type BackoffFunc func(retry int) (sleep time.Duration)

func newDefaultBackoffPolicy(mind, maxd time.Duration) BackoffFunc {

	factor := time.Second
	for _, scale := range []time.Duration{time.Hour, time.Minute, time.Second, time.Millisecond, time.Microsecond, time.Nanosecond} {
		d := mind.Truncate(scale)
		if d > 0 {
			factor = scale
			break
		}
	}

	return func(retry int) (sleep time.Duration) {
		wait := 2 << max(0, min(32, retry)) * factor
		jitter := time.Duration(rand.Int64N(int64(max(1, int(wait)/5)))) // max 20% jitter
		wait = mind + wait + jitter
		if wait > maxd {
			wait = maxd
		}
		return wait
	}
}
