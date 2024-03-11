package gotracingutil

import "context"

func Parse(ctx context.Context) (labels map[string]string) {
	if ctx == nil {
		return nil
	}

	labels = make(map[string]string)
	for _, key := range tracingKeys {
		value := ctx.Value(key)
		if value == nil {
			labels[key] = ""
		} else {
			labels[key] = value.(string)
		}
	}
	return labels
}
