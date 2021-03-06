package util

import (
	"context"
	"time"

	"github.com/juju/errors"
	"github.com/ngaut/log"
)

// RetryOnError try a possible fail event several times before fail
func RetryOnError(ctx context.Context, retryCount int, fn func() error) error {
	var err error
	for i := 0; i < retryCount; i++ {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		err = fn()
		if err == nil {
			break
		}

		log.Errorf("retry error %d times, %v", i+1, err)
		Sleep(ctx, 2*time.Second)
	}

	return errors.Trace(err)
}

// Sleep defines special `sleep` with context
func Sleep(ctx context.Context, sleepTime time.Duration) {
	ticker := time.NewTicker(sleepTime)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return
	case <-ticker.C:
		return
	}
}

// TimeoutContext create context with timeout
func TimeoutContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}
