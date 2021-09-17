package github.com/ybm2dyd/galog

import (
	"fmt"
	"os"
	"path"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

//FileHandler writes log to a file.
type FileHandler struct {
	fd *os.File
}

// NewFileHandler new filehandler
func NewFileHandler(fileName string, flag int) (*FileHandler, error) {
	dir := path.Dir(fileName)
	os.Mkdir(dir, 0777)

	f, err := os.OpenFile(fileName, flag, 0)
	if err != nil {
		return nil, err
	}

	h := new(FileHandler)

	h.fd = f

	return h, nil
}

func (h *FileHandler) Write(b []byte) (n int, err error) {
	return h.fd.Write(b)
}

// Close file handler
func (h *FileHandler) Close() error {
	if h.fd != nil {
		return h.fd.Close()
	}
	return nil
}

//RotatingFileHandler writes log a file, if file size exceeds maxBytes,
//it will backup current file and open a new one.
//
//max backup file number is set by backupCount, it will delete oldest if backups too many.
type RotatingFileHandler struct {
	fd *os.File

	fileName    string
	maxBytes    int
	curBytes    int
	backupCount int
	curCount    int
	mutex       sync.Mutex
}

// NewRotatingFileHandler return RotatingFileHandler
func NewRotatingFileHandler(fileName string, maxBytes int, backupCount int) (*RotatingFileHandler, error) {
	dir := path.Dir(fileName)
	os.MkdirAll(dir, 0777)

	h := new(RotatingFileHandler)

	if maxBytes <= 0 {
		return nil, fmt.Errorf("invalid max bytes")
	}

	h.fileName = fileName
	h.maxBytes = maxBytes
	h.backupCount = backupCount

	var err error
	h.curCount = 0
	h.fd, err = os.OpenFile(h.fileName+".0", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}

	f, err := h.fd.Stat()
	if err != nil {
		return nil, err
	}
	h.curBytes = int(f.Size())

	return h, nil
}

func (h *RotatingFileHandler) Write(p []byte) (n int, err error) {
	if h.backupCount > 0 && h.shouldRollover(len(p)) {
		h.mutex.Lock()
		err = h.doRollover()
		h.mutex.Unlock()
		if err != nil {
			return 0, err
		}
	}

	n, err = h.fd.Write(p)
	h.curBytes += n
	return
}

// Close file handler
func (h *RotatingFileHandler) Close() error {
	if h.fd != nil {
		return h.fd.Close()
	}
	return nil
}

func (h *RotatingFileHandler) shouldRollover(length int) bool {
	if h.maxBytes > 0 && h.curBytes+length >= h.maxBytes {
		return true
	}
	return false
}

func (h *RotatingFileHandler) doRollover() error {

	if h.curBytes < h.maxBytes {
		return nil
	}

	h.curCount = (h.curCount + 1) % h.backupCount
	dfn := fmt.Sprintf("%s.%d", h.fileName, h.curCount)
	newFd, err := os.OpenFile(dfn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	oldFd := atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&h.fd)), unsafe.Pointer(newFd))
	(*os.File)(oldFd).Close()
	h.curBytes = 0
	return nil
}

//TimeRotatingFileHandler writes log to a file,
//it will backup current and open a new one, with a period time you sepecified.
//
//refer: http://docs.python.org/2/library/logging.handlers.html.
//same like python TimedRotatingFileHandler.
type TimeRotatingFileHandler struct {
	fd *os.File

	baseName   string
	when       Rollingtime
	interval   int64
	suffix     string
	rolloverAt int64
	mutex      sync.Mutex
}

type Rollingtime int

const (
	// WhenSecond rotated by second
	WhenSecond Rollingtime = iota
	// WhenMinute rotated by by minute
	WhenMinute
	// WhenHour rotated by by hour
	WhenHour
	// WhenDay rotated by by Day
	WhenDay
)

func rolloverAt(now time.Time, h *TimeRotatingFileHandler) error {
	switch h.when {
	case WhenSecond:
		h.rolloverAt = now.Unix() + int64(h.interval)
	case WhenMinute:
		h.rolloverAt = (time.Date(now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), 0, 0, time.Local)).Unix() + int64(h.interval)*60
	case WhenHour:
		h.rolloverAt = (time.Date(now.Year(), now.Month(), now.Day(),
			now.Hour(), 0, 0, 0, time.Local)).Unix() + int64(h.interval)*3600
	case WhenDay:
		h.rolloverAt = (time.Date(now.Year(), now.Month(), now.Day(),
			0, 0, 0, 0, time.Local)).Unix() + int64(h.interval)*3600*24
	default:
		return fmt.Errorf("invalid when_rotate: %d", h.when)
	}
	return nil
}

// NewTimeRotatingFileHandler return TimeRotatingFileHandler
func NewTimeRotatingFileHandler(baseName string, when Rollingtime, interval int) (*TimeRotatingFileHandler, error) {
	dir := path.Dir(baseName)
	os.Mkdir(dir, 0777)

	h := new(TimeRotatingFileHandler)

	h.baseName = baseName
	h.interval = int64(interval)
	h.when = when

	now := time.Now()
	err := rolloverAt(now, h)
	if err != nil {
		return nil, err
	}

	switch h.when {
	case WhenSecond:
		h.suffix = "2006-01-02_15-04-05"
	case WhenMinute:
		h.suffix = "2006-01-02_15-04"
	case WhenHour:
		h.suffix = "2006-01-02_15"
	case WhenDay:
		h.suffix = "2006-01-02"
	default:
		return nil, fmt.Errorf("invalid when_rotate: %d", h.when)
	}

	h.fd, err = open(h.baseName + "." + now.Format(h.suffix))
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *TimeRotatingFileHandler) doRollover() error {
	//refer http://hg.python.org/cpython/file/2.7/Lib/logging/handlers.py
	now := time.Now()
	if h.rolloverAt > now.Unix() {
		return nil
	}

	fName := h.baseName + "." + now.Format(h.suffix)
	newFd, err := open(fName)
	if err != nil {
		return err
	}
	oldFd := atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&h.fd)), unsafe.Pointer(newFd))
	err = rolloverAt(now, h)
	if err != nil {
		return err
	}
	(*os.File)(oldFd).Close()
	return err
}

func (h *TimeRotatingFileHandler) shouldRollover() bool {
	now := time.Now()
	if h.rolloverAt <= now.Unix() {
		return true
	}
	return false
}

func open(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

func (h *TimeRotatingFileHandler) Write(b []byte) (n int, err error) {
	if h.shouldRollover() {
		h.mutex.Lock()
		err = h.doRollover()
		h.mutex.Unlock()
		if err != nil {
			return 0, err
		}
	}

	return h.fd.Write(b)
}

// Close file handler
func (h *TimeRotatingFileHandler) Close() error {
	return h.fd.Close()
}
