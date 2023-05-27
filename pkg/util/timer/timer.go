package timer

import (
	"time"

	"github.com/Lyfebloc/logger"
)

func Start(task any) func() {
	start := time.Now()

	return func() {
		logger.Infof("processed task: %v in: %v", task, time.Since(start))
	}
}
