package trace_context

import (
	"context"
	"github.com/google/uuid"
)

const (
	TraceID      = "Trace-ID"
	SpanID       = "Span-ID"
	ParentSpanID = "Parent-Span-ID"
)

func NewContext(ctx context.Context) context.Context {
	traceID, ok := ctx.Value(TraceID).(string)
	if !ok {
		traceID = uuid.NewString()
	}
	newCtx := context.WithValue(ctx, TraceID, traceID)

	spanID := uuid.NewString()
	newCtx = context.WithValue(newCtx, SpanID, spanID)

	parentCtxSpanID, ok := ctx.Value(SpanID).(string)
	if !ok {
		newCtx = context.WithValue(newCtx, ParentSpanID, parentCtxSpanID)
	}

	return newCtx
}
