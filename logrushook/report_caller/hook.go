package report_caller

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type LocationHandler func(fileAbsolutePath string, line int) string

var defaultLocationHandler LocationHandler = func(fileAbsolutePath string, line int) string {
	return fmt.Sprintf("%s:%d", fileAbsolutePath, line)
}

type Hook struct {
	levels          []logrus.Level
	fieldKey        string
	locationHandler LocationHandler
}

func NewHook(levels []logrus.Level, fieldKey string) *Hook {
	return &Hook{
		levels:          levels,
		fieldKey:        fieldKey,
		locationHandler: defaultLocationHandler,
	}
}

func (hook *Hook) SetLocationHandler(handler LocationHandler) {
	hook.locationHandler = handler
}

func (hook *Hook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	caller := getCaller()
	entry.Data[hook.fieldKey] = hook.locationHandler(caller.File, caller.Line)
	return nil
}

// copied from https://github.com/sirupsen/logrus

var (
	// qualified package name, cached at first use
	logrusPackage string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
)

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

// getCaller retrieves the name of the first non-logrus calling function
func getCaller() *runtime.Frame {
	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()

			if strings.Contains(funcName, "fireHooks") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != logrusPackage {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}
