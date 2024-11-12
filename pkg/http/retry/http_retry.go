package retry

import (
	"errors"
	"math/rand"
	"time"
)

type Retryer struct {
}

func (r *Retryer) Retry() {

}

func Retry(retryTimes int, requestTimeOut time.Duration, requestFn func() error) error {
	var err error
	for i := 0; i < retryTimes; i++ {
		err = requestFn()
		if err == nil {
			return nil
		}
		if i == retryTimes {
			break
		}
		// Calculate exponential backoff with random salt
		delay := time.Second*time.Duration(1<<i) + time.Duration(rand.Intn(1000))*time.Millisecond
		time.Sleep(delay)
	}
	msg := "all retry attempts failed: "
	if err != nil {
		msg += err.Error()
	}
	return errors.New(msg)
}
