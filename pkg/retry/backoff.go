package retry

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

// GetExponentialBackoff - Configure exponential backoff options
func GetExponentialBackoff() *backoff.ExponentialBackOff {

	expBackoff := backoff.NewExponentialBackOff()

	// MaxElaspedTime - Retry mechanism will keep retrying for time set
	expBackoff.MaxElapsedTime = 15 * time.Second

	// MaxInterval - Max time between successive retries
	expBackoff.MaxInterval = 4 * time.Second

	// Add random seeding
	expBackoff.Multiplier = 2

	return expBackoff
}
