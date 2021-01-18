package utils

import (
	"errors"
	"log"
	"time"
)

const (
	maxRetries = 5
)

/*
 * Retry - Retries the function block that is passed as a parameter
 * @timeOut - default wait time
 * @incremental - Incremental timeOut for every iteration
 * @callback - function to execute on retry
 */
func Retry(timeOut, incremental int, callback func() error) (err error) {
	attempts := 0
	prevTime := 1
	for {
		attempts++
		err = callback()
		if err == nil {
			return err
		}
		if attempts == maxRetries {
			return errors.New("Max attempts are reached")
		}
		time.Sleep(time.Duration(prevTime*timeOut) * time.Second)
		log.Printf("Retrying attempt#%d after waiting %ds with last error: %s\n", attempts, prevTime*timeOut, err)
		prevTime = incremental * prevTime
	}
}
