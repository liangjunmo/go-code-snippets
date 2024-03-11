package gotracingutil

func init() {
	resetTracingKeys()
}

const (
	defaultTracingIDKey = "TracingID"
)

var (
	tracingIDKey string
	tracingKeys  []string
)

func resetTracingKeys() {
	tracingIDKey = defaultTracingIDKey
	tracingKeys = []string{tracingIDKey}
}

func SetTracingIDKey(key string) {
	tracingIDKey = key
	tracingKeys[0] = tracingIDKey
}

func AppendTracingKeys(keys []string) {
	tracingKeys = append(tracingKeys, keys...)
}
