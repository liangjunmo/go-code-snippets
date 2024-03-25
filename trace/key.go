package trace

func init() {
	resetTraceKeys()
}

const (
	defaultTraceIDKey = "TraceID"
)

var (
	traceIDKey string
	traceKeys  []string
)

func resetTraceKeys() {
	traceIDKey = defaultTraceIDKey
	traceKeys = []string{traceIDKey}
}

func SetTraceIDKey(key string) {
	traceIDKey = key
	traceKeys[0] = traceIDKey
}

func AppendTraceKeys(keys []string) {
	traceKeys = append(traceKeys, keys...)
}
