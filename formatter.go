package github.com/ybm2dyd/galog

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	defaultTimestampFormat = time.RFC3339
	callDepth              = 5
	red                    = 31
	yellow                 = 33
	purple                 = 34
	blue                   = 36
	gray                   = 37
)

// Formatter Format is expected to return an array of bytes which are then
// logged to `logger.Out`.
type Formatter interface {
	Format(level Level, buffer *bytes.Buffer, msg string) ([]byte, error)
}

// TextFormatter formats logs into text
type TextFormatter struct {
	// Force disabling colors.
	DisableColors bool

	// Disable timestamp logging. useful when output is redirected to logging
	// system that already adds timestamps.
	DisableTimestamp bool

	// TimestampFormat to use for display when a full timestamp is printed
	TimestampFormat string

	DisableFormat bool
}

// Format renders a single log entry
func (f *TextFormatter) Format(level Level, buffer *bytes.Buffer, msg string) ([]byte, error) {
	if f.DisableFormat {
		return []byte(msg), nil
	} else {
		var levelColor int
		switch level {
		case DebugLevel, TraceLevel:
			levelColor = gray
		case WarnLevel:
			levelColor = yellow
		case ErrorLevel, FatalLevel, PanicLevel:
			levelColor = red
		case InfoLevel:
			levelColor = blue
		default:
			levelColor = blue
		}

		levelText := strings.ToUpper(level.String())

		timestampFormat := f.TimestampFormat
		if timestampFormat == "" {
			timestampFormat = defaultTimestampFormat
		}

		if !f.DisableColors {
			fmt.Fprintf(buffer, "\x1b[%dm", levelColor)
		}
		if !f.DisableTimestamp {
			fmt.Fprintf(buffer, "%-5s [%s]", levelText, time.Now().Format(timestampFormat))
		} else {
			fmt.Fprintf(buffer, "%-5s", levelText)
		}

		pc, file, line, ok := runtime.Caller(callDepth)
		pcName := "???"
		if !ok {
			file = "???"
			line = 0
		} else {
			pcName = runtime.FuncForPC(pc).Name()
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					file = file[i+1:]
					break
				}
			}
		}
		buffer.WriteString(" " + file + ":" + strconv.FormatInt(int64(line), 10) + " " + pcName)

		if !f.DisableColors {
			buffer.WriteString("\x1b[0m")
		}
		buffer.WriteString(" " + msg)
	}

	return buffer.Bytes(), nil
}
