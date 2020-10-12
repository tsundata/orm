package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\u001b[31m[error]\u001b[0m ", log.LstdFlags|log.Lshortfile)
	warnLog  = log.New(os.Stdout, "\u001b[33m[warn]\u001b[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\u001b[35m[info]\u001b[0m ", log.LstdFlags|log.Lshortfile)
	debugLog = log.New(os.Stdout, "\u001b[37m[debug]\u001b[0m ", log.LstdFlags|log.Lshortfile)
	traceLog = log.New(os.Stdout, "\u001b[37m[trace]\u001b[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Warn   = warnLog.Println
	Warnf  = warnLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
	Debug  = debugLog.Println
	Debugf = debugLog.Printf
	Trace  = traceLog.Println
	Tracef = traceLog.Printf
)

const (
	TraceLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	Disabled
)

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if WarnLevel < level {
		warnLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
	if DebugLevel < level {
		debugLog.SetOutput(ioutil.Discard)
	}
	if TraceLevel < level {
		traceLog.SetOutput(ioutil.Discard)
	}
}
