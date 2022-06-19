package logpush

import (
	"sync"
	"time"
)

var (
	pool          []string
	lastFlushTime = time.Now()
	poolLength    = 10000
)

type Engine interface {
	Flush([]string) error
}

type LogPush struct {
	MaxPoolLength int
	PushDuration  time.Duration
	Engine        Engine
	mux           sync.Mutex
}

// Push 根据初始化LogPush条件，决定是否要推送日志。
// Decide whether to Push logs based on the initial Log Push condition.
func (l *LogPush) Push(log string) error {
	l.mux.Lock()
	defer l.mux.Unlock()

	if l.MaxPoolLength == 0 {
		l.MaxPoolLength = 10000
	}

	if l.PushDuration == time.Duration(0) {
		l.PushDuration = time.Minute * 5
	}

	pool = append(pool, log)
	if len(pool) >= l.MaxPoolLength || lastFlushTime.Add(l.PushDuration).Before(time.Now()) {
		lastFlushTime = time.Now()
		err := l.Engine.Flush(pool)
		if err != nil {
			if len(pool) > poolLength {
				pool = []string{}
			}
			return err
		}
		pool = []string{}
	}
	return nil
}

// Flush 立刻将日志池所有日志推送
// All logs in the log pool are pushed immediately
func (l *LogPush) Flush() error {
	l.mux.Lock()
	defer l.mux.Unlock()

	err := l.Engine.Flush(pool)
	if err != nil {
		if len(pool) > poolLength {
			pool = []string{}
		}
		return err
	}
	pool = []string{}
	return nil
}
