package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

const parentSpan = "PARENT_SPAN"

func CtxWithParentSpan(ctx context.Context, parent opentracing.Span) context.Context {
	return context.WithValue(ctx, parentSpan, parent)
}

func StartChildSpan(ctx context.Context, spanName string) opentracing.Span {
	val := ctx.Value(parentSpan)
	if nil != val {
		sp := val.(opentracing.Span)
		return opentracing.StartSpan(spanName, opentracing.ChildOf(sp.Context()))
	}
	return opentracing.StartSpan(spanName)
}
