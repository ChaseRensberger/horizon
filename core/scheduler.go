package core

import (
	"time"
)

func happenEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}
