package function

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Retry calls f() until successful, if it fails then it waits
// for a period that starts with 2 seconds and increases exponentially
// until it is capped at 1 minute.
func Retry(f func() error, interrupt chan os.Signal) {
	var sleep time.Duration = 2
	var retryTimestamp time.Time = time.Now()

	for {
		select {
		case <-interrupt:
			log.Warn("Interrupt/kill signal received, quitting.")
			os.Exit(1)
		default:
			// call f() if the current time passed retryTimestamp
			if time.Now().After(retryTimestamp) {
				err := f()
				if err != nil {
					log.Error(err)

					// couldn't connect, retry.
					log.Infof("Retrying in %d seconds", sleep)
					retryTimestamp = time.Now().Add(sleep * time.Second)

					// increase exponentially, cap at ~ 1 minute (64 seconds).
					if sleep < 64 {
						sleep = sleep * 2
					}
				} else {
					return
				}
			}

			// loop every second, check for signals and timeline in each iteration.
			time.Sleep(1 * time.Second)
		}
	}
}
