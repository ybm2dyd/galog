package galog

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

// Level type
type Level uint32

// These are the different logging levels. You can set the logging level to log
// on your instance of logger
const (
	// PanicLevel level, highest level of severity.
	PanicLevel Level = iota
	// FatalLevel level.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	case "trace":
		return TraceLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid galog Level: %q", lvl)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (level *Level) UnmarshalText(text []byte) error {
	l, err := ParseLevel(string(text))
	if err != nil {
		return err
	}

	*level = l

	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case TraceLevel:
		return []byte("trace"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case InfoLevel:
		return []byte("info"), nil
	case WarnLevel:
		return []byte("warn"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case PanicLevel:
		return []byte("panic"), nil
	}

	return nil, fmt.Errorf("not a valid galog level %d", level)
}

// Convert the Level to a string. E.g. PanicLevel becomes "panic".
func (level Level) String() string {
	if b, err := level.MarshalText(); err == nil {
		return string(b)
	}
	return "unknown"
}

type Logger struct {
	Level     Level
	Out       Handler
	Formatter Formatter
	mu        MutexWrap
	Buffer    *bytes.Buffer
}

type MutexWrap struct {
	lock     sync.Mutex
	disabled bool
}

func (mw *MutexWrap) Lock() {
	if !mw.disabled {
		mw.lock.Lock()
	}
}

func (mw *MutexWrap) Unlock() {
	if !mw.disabled {
		mw.lock.Unlock()
	}
}

func (mw *MutexWrap) Disable() {
	mw.disabled = true
}

// New It's recommended to make this a global instance called `log`.
func New() *Logger {
	return &Logger{
		Out:       os.Stdout,
		Formatter: new(TextFormatter),
		Level:     InfoLevel,
	}
}

func (logger *Logger) Log(level Level, args ...interface{}) error {
	if logger.IsLevelEnabled(level) {
		return logger.log(level, fmt.Sprint(args...))
	}
	return nil
}

func (logger *Logger) Trace(args ...interface{}) error {
	return logger.Log(TraceLevel, args...)
}

func (logger *Logger) Debug(args ...interface{}) error {
	return logger.Log(DebugLevel, args...)
}

func (logger *Logger) Print(args ...interface{}) error {
	return logger.Info(args...)
}

func (logger *Logger) Info(args ...interface{}) error {
	return logger.Log(InfoLevel, args...)
}

func (logger *Logger) Warn(args ...interface{}) error {
	return logger.Log(WarnLevel, args...)
}

func (logger *Logger) Warning(args ...interface{}) error {
	return logger.Warn(args...)
}

func (logger *Logger) Error(args ...interface{}) error {
	return logger.Log(ErrorLevel, args...)
}

func (logger *Logger) Fatal(args ...interface{}) error {
	return logger.Log(FatalLevel, args...)
}

func (logger *Logger) Panic(args ...interface{}) error {
	return logger.Log(PanicLevel, args...)
}

func (logger *Logger) Logf(level Level, format string, args ...interface{}) error {
	if logger.IsLevelEnabled(level) {
		return logger.Log(level, fmt.Sprintf(format, args...))
	}
	return nil
}

func (logger *Logger) Tracef(format string, args ...interface{}) error {
	return logger.Logf(TraceLevel, format, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) error {
	return logger.Logf(DebugLevel, format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) error {
	return logger.Logf(InfoLevel, format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) error {
	return logger.Infof(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) error {
	return logger.Logf(WarnLevel, format, args...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) error {
	return logger.Warnf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) error {
	return logger.Logf(ErrorLevel, format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) error {
	return logger.Logf(FatalLevel, format, args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) error {
	return logger.Logf(PanicLevel, format, args...)
}

func (logger *Logger) Logln(level Level, args ...interface{}) error {
	if logger.IsLevelEnabled(level) {
		return logger.Log(level, fmt.Sprintln(args...))
	}
	return nil
}

func (logger *Logger) Traceln(args ...interface{}) error {
	return logger.Logln(TraceLevel, args...)
}

func (logger *Logger) Debugln(args ...interface{}) error {
	return logger.Logln(DebugLevel, args...)
}

func (logger *Logger) Infoln(args ...interface{}) error {
	return logger.Logln(InfoLevel, args...)
}

func (logger *Logger) Println(args ...interface{}) error {
	return logger.Infoln(args...)
}

func (logger *Logger) Warnln(args ...interface{}) error {
	return logger.Logln(WarnLevel, args...)
}

func (logger *Logger) Warningln(args ...interface{}) error {
	return logger.Warnln(args...)
}

func (logger *Logger) Errorln(args ...interface{}) error {
	return logger.Logln(ErrorLevel, args...)
}

func (logger *Logger) Fatalln(args ...interface{}) error {
	return logger.Logln(FatalLevel, args...)
}

func (logger *Logger) Panicln(args ...interface{}) error {
	return logger.Logln(PanicLevel, args...)
}

func (logger *Logger) log(level Level, msg string) error {
	var buffer *bytes.Buffer

	buffer = getBuffer()
	defer func() {
		putBuffer(buffer)
	}()
	buffer.Reset()

	serialized, err := logger.Formatter.Format(level, buffer, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		return err
	}
	logger.mu.Lock()
	defer logger.mu.Unlock()
	if _, err := logger.Out.Write(serialized); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
	}
	return err
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func (logger *Logger) SetNoLock() {
	logger.mu.Disable()
}

func (logger *Logger) level() Level {
	return Level(atomic.LoadUint32((*uint32)(&logger.Level)))
}

// SetLevel sets the logger level.
func (logger *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&logger.Level), uint32(level))
}

// GetLevel returns the logger level.
func (logger *Logger) GetLevel() Level {
	return logger.level()
}

func (logger *Logger) IsLevelEnabled(level Level) bool {
	return logger.level() >= level
}

// SetFormatter sets the logger formatter.
func (logger *Logger) SetFormatter(formatter Formatter) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.Formatter = formatter
}

// SetOutput sets the logger output.
func (logger *Logger) SetOutput(output Handler) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.Out = output
}
