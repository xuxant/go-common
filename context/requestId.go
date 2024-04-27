package context

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type requestIdKey struct{}

// WithRequestId return new context with request id
func WithRequestid(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIdKey{}, id)
}

func WithGenerateRequestId(ctx context.Context) context.Context {
	id := uuid.New().String()
	ctx = context.WithValue(ctx, requestIdKey{}, id)
	ctx = SetLoggerRequestID(ctx, id)
	return ctx
}

func SetLoggerRequestID(ctx context.Context, id string) context.Context {
	return log.With().Str("request-id", id).Logger().WithContext(ctx)
}

// RequestIdFromContext returns request id from context
func RequestIdFromContext(ctx context.Context) string {
	if id, ok := ctx.Value(requestIdKey{}).(string); ok {
		return id
	}
	return ""
}
