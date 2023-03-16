package weaverx

import (
	"errors"
	"time"

	"github.com/ServiceWeaver/weaver"
	"github.com/berquerant/weaver-pokemon-type/errorx"
)

//go:generate go run github.com/berquerant/goconfig@v0.2.0 -prefix Retry -option -output retry_config_generated.go -field "MaxAttempts int|MaxBackoff time.Duration|InitialBackoff time.Duration"

var (
	ErrRetryExhausted = errors.New("RetryExhausted")
)

func Retry[T any](
	f func() (T, error),
	opt ...RetryConfigOption,
) (T, error) {
	config := NewRetryConfigBuilder().
		MaxAttempts(5).
		MaxBackoff(time.Second * 3).
		InitialBackoff(time.Millisecond * 300).
		Build()
	config.Apply(opt...)

	backoff := config.InitialBackoff.Get()
	for i := 0; i < config.MaxAttempts.Get(); i++ {
		value, err := f()

		if errors.Is(err, weaver.ErrRetriable) {
			time.Sleep(backoff)
			backoff *= 2
			if backoff > config.MaxBackoff.Get() {
				backoff = config.MaxBackoff.Get()
			}
			continue
		}

		return value, nil
	}

	var v T
	return v, errorx.New(ErrRetryExhausted)
}
