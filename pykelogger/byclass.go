package pykelogger

import (
	"fmt"
	"os"
)

func NewLogger(path string) (*Logger, error) {
	l := &Logger{logPath: path}
	logFile, fileErr := os.OpenFile(l.logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if fileErr != nil {
		l.logFile = nil
		return l, fileErr
	}

	l.logFile = logFile

	return l, nil
}

func (l *Logger) Log(format string, args ...any) {
	timeStamp := getTimeStamp()
	if l.logFile != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		message := fmt.Sprintf(format, args...)
		logLine := fmt.Sprintf("%s : %s\r\n", timeStamp, message)
		l.logFile.WriteString(logLine)
		l.logFile.Sync()
	} else {
		prefix := fmt.Sprintf("%s[ERROR] %s ", colorRed, timeStamp)
		message := fmt.Sprintf(format, args...)
		fmt.Printf("%s%s%s\n", prefix, message, colorReset)
	}
}

func (l *Logger) Infof(format string, args ...any) {
	timeStamp := getTimeStamp()
	if l.logFile != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		message := fmt.Sprintf(format, args...)
		logLine := fmt.Sprintf("[INFO] %s : %s\r\n", timeStamp, message)
		l.logFile.WriteString(logLine)
		l.logFile.Sync()
	} else {
		prefix := fmt.Sprintf("%s[INFO] %s ", colorGreen, timeStamp)
		message := fmt.Sprintf(format, args...)
		fmt.Printf("%s%s%s\n", prefix, message, colorReset)
	}
}

func (l *Logger) Errorf(format string, args ...any) {
	timeStamp := getTimeStamp()
	if l.logFile != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		message := fmt.Sprintf(format, args...)
		logLine := fmt.Sprintf("[ERROR] %s : %s\r\n", timeStamp, message)
		l.logFile.WriteString(logLine)
		l.logFile.Sync()
	} else {
		prefix := fmt.Sprintf("%s[ERROR] %s ", colorRed, timeStamp)
		message := fmt.Sprintf(format, args...)
		fmt.Printf("%s%s%s\n", prefix, message, colorReset)
	}
}

func (l *Logger) Tokenf(format string, args ...any) {
	if l.logFile != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		logLine := fmt.Sprintf(format, args...)
		l.logFile.WriteString(logLine)
		l.logFile.Sync()
	} else {
		message := fmt.Sprintf(format, args...)
		fmt.Printf("%s%s%s", colorBrown, message, colorReset)
	}
}

func (l *Logger) Fatalf(format string, args ...any) {
	timeStamp := getTimeStamp()
	if l.logFile != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		message := fmt.Sprintf(format, args...)
		logLine := fmt.Sprintf("[FATAL] %s : %s\r\n", timeStamp, message)
		l.logFile.WriteString(logLine)
		l.logFile.Sync()
	} else {
		prefix := fmt.Sprintf("%s[FATAL] %s ", colorRed, timeStamp)
		message := fmt.Sprintf(format, args...)
		fmt.Printf("%s%s%s\n", prefix, message, colorReset)
	}
	os.Exit(1)
}

func (l *Logger) Close() {
	l.logFile.Close()
}
