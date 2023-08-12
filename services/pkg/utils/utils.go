package utils

import (
	"fmt"
	"time"
)

type RetryFunc[T any] func(address string) (T, error)

// RetryConnect attempts to execute the given function and retries to connect to the third party  for a maximum number of attempts.
// It waits for the specified delay between each retry.
func RetryConnect[T any](attempts uint, delay time.Duration, address string, fn RetryFunc[T]) (T, error) {
	var connection T
	var err error
	for i := 0; i < int(attempts); i++ {
		connection, err = fn(address)
		if err == nil {
			return connection, nil
		}
		seconds := delay.Seconds()
		fmt.Printf("we cant connect to %s so we retry it after %.1f secconds  : %s \n", address, seconds, err)
		time.Sleep(delay)
	}
	return connection, err
}
