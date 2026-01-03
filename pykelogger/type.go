package pykelogger

import (
	"os"
	"sync"
)

type (
	Logger struct {
		logPath string
		logFile *os.File
		mu      sync.Mutex
	}
)
