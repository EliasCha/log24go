package log24go

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
)

type Log24go struct {
	Out io.Writer
	Mu  sync.Mutex
	Buf []byte
}

func (log *Log24go) Writer(level string, s string) error {
	tn := time.Now()

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "N/A"
		line = 0
	}

	log.Mu.Lock()
	defer log.Mu.Unlock()

	// Set format for log message
	log.WriterFormatter(level, file, line, tn)

	// Append the log message s
	log.Buf = append(log.Buf, s...)

	// Set newline for end of log line
	log.Buf = append(log.Buf, "\n"...)

	_, err := log.Out.Write(log.Buf)
	return err
}

func (log *Log24go) WriterFormatter(level string, file string, line int, tn time.Time) {
	// Zero the buffer
	log.Buf = log.Buf[:0]

	// Create variable for shorter name
	var buf *[]byte = &log.Buf

	*buf = append(*buf, fmt.Sprintf("[%s]", level)...)
	*buf = append(*buf, " "...)

	// Set date format
	var year int
	var month time.Month
	var date int
	year, month, date = tn.Date()
	*buf = append(*buf, fmt.Sprintf("%d-%d-%d", year, month, date)...)
	*buf = append(*buf, " "...)

	// Set time format
	var hour, minute, second, nanosecond int
	var err error

	if tn.Hour() < 10 {
		hour, err = strconv.Atoi(fmt.Sprintf("0%d", tn.Hour()))
		if err != nil {
			hour = tn.Hour()
		}
	} else {
		hour = tn.Hour()
	}

	if tn.Minute() < 10 {
		minute, err = strconv.Atoi(fmt.Sprintf("0%d", tn.Minute()))
		if err != nil {
			minute = tn.Minute()
		}
	} else {
		minute = tn.Minute()
	}

	if tn.Second() < 10 {
		second, err = strconv.Atoi(fmt.Sprintf("0%d", tn.Second()))
		if err != nil {
			second = tn.Second()
		}
	} else {
		second = tn.Second()
	}

	nanosecond = tn.Nanosecond()
	*buf = append(*buf, fmt.Sprintf("%d:%d:%d.%d", hour, minute, second, nanosecond)...)
	*buf = append(*buf, " "...)

	// Set file and file line from where it was called
	*buf = append(*buf, fmt.Sprintf("%s:%d:", file, line)...)
	*buf = append(*buf, " "...)
}

func (log *Log24go) Info(inter ...interface{}) {
	log.Writer(INFO, fmt.Sprint(inter...))
}

func (log *Log24go) Warn(inter ...interface{}) {
	log.Writer(WARN, fmt.Sprint(inter...))
}

func (log *Log24go) Error(inter ...interface{}) {
	log.Writer(ERROR, fmt.Sprint(inter...))
}

func (log *Log24go) Debug(inter ...interface{}) {
	log.Writer(DEBUG, fmt.Sprint(inter...))
}

func NewLog24go(out io.Writer) *Log24go {
	return &Log24go{Out: out}
}
