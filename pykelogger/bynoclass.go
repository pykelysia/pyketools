package pykelogger

import "os"

func UseFile(path string) error {
	L.logPath = path
	logFile, fileErr := os.OpenFile(L.logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if fileErr != nil {
		L.logFile = nil
		return fileErr
	}
	L.logFile = logFile
	return nil
}

func CloseFile() {
	L.logFile.Close()
}
