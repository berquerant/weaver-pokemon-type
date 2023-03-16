package httpx

import (
	"context"
	"net/http"

	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type weaverInstanceKeyType string

var weaverInstanceKey weaverInstanceKeyType = "weaverInstanceKey"

func WeaverInstanceWithContext(ctx context.Context, instance weaver.Instance) context.Context {
	return context.WithValue(ctx, weaverInstanceKey, instance)
}

func WeaverInstanceFromContext(ctx context.Context) weaver.Instance {
	return ctx.Value(weaverInstanceKey).(weaver.Instance)
}

func InjectWeaverInstance(instance weaver.Instance, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r.WithContext(
			WeaverInstanceWithContext(
				requestIDWithContext(r.Context()), instance)))
	}
}

type requestIDKeyType string

var requestIDKey requestIDKeyType = "requestIDKey"

func requestIDWithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, requestIDKey, uuid.New().String())
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}
