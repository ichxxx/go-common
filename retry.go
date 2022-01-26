package common

import (
	"time"
)

const defaultRetryTimes = 3

type retry struct {
	fn    func() error
	times int
	wait  time.Duration
}

func Retry(fn func() error) retry {
	return retry{
		times: defaultRetryTimes,
		fn:    fn,
	}
}

func (r retry) Times(t int) retry {
	r.times = t
	return r
}

func (r retry) Wait(d time.Duration) retry {
	r.wait = d
	return r
}

func (r retry) Do() error {
	var err error
	for i := 0; i < r.times; i++ {
		err = r.fn()
		if err == nil {
			return nil
		}
		if r.wait > 0 {
			time.Sleep(r.wait)
		}
	}
	return err
}

