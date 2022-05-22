package opentracingbridge

import (
	"context"

	"go.opentelemetry.io/otel/bridge/opentracing"
	"go.opentelemetry.io/otel/bridge/opentracing/migration"
	"go.opentelemetry.io/otel/trace"
)

// WarningHandler is a type of handler that receives warnings
// from the BridgeTracer.
type WarningHandler = opentracing.BridgeWarningHandler

// Tracer is an implementation of the OpenTracing tracer, which
// translates the calls to the OpenTracing API into OpenTelemetry
// counterparts and calls the underlying OpenTelemetry tracer.
type Tracer = opentracing.BridgeTracer

// NewBridgeTracer creates a new BridgeTracer. The new tracer forwards
// the calls to the OpenTelemetry Noop tracer, so it should be
// overridden with the SetOpenTelemetryTracer function. The warnings
// handler does nothing by default, so to override it use the
// SetWarningHandler function.
func NewBridgeTracer() *Tracer {
	return opentracing.NewBridgeTracer()
}

// NewTracerPair is a utility function that creates a BridgeTracer and a
// WrapperTracerProvider. WrapperTracerProvider creates a single instance of
// WrapperTracer. The BridgeTracer forwards the calls to the WrapperTracer
// that wraps the passed tracer. BridgeTracer and WrapperTracerProvider are
// returned to the caller and the caller is expected to register BridgeTracer
// with opentracing and WrapperTracerProvider with opentelemetry.
func NewTracerPair(tracer trace.Tracer) (*opentracing.BridgeTracer, *WrapperTracerProvider) {
	return opentracing.NewTracerPair(tracer)
}

func NewTracerPairWithContext(ctx context.Context, tracer trace.Tracer) (context.Context, *Tracer, *WrapperTracerProvider) {
	return opentracing.NewTracerPairWithContext(ctx, tracer)
}

type WrapperTracerProvider = opentracing.WrapperTracerProvider

// NewWrappedTracerProvider creates a new trace provider that creates a single
// instance of WrapperTracer that wraps OpenTelemetry tracer.
func NewWrappedTracerProvider(bridge *Tracer, tracer trace.Tracer) *WrapperTracerProvider {
	return opentracing.NewWrappedTracerProvider(bridge, tracer)
}

// WrapperTracer is a wrapper around an OpenTelemetry tracer. It
// mostly forwards the calls to the wrapped tracer, but also does some
// extra steps like setting up a context with the active OpenTracing
// span.
//
// It does not need to be used when the OpenTelemetry tracer is also
// aware how to operate in environment where OpenTracing API is also
// used.
type WrapperTracer = opentracing.WrapperTracer

// NewWrapperTracer wraps the passed tracer and also talks to the
// passed bridge tracer when setting up the context with the new
// active OpenTracing span.
func NewWrapperTracer(bridge *Tracer, tracer trace.Tracer) *WrapperTracer {
	return opentracing.NewWrapperTracer(bridge, tracer)
}

// MigrationDeferredContextSetupTracerExtension is an interface an
// OpenTelemetry tracer may implement in order to cooperate with the
// calls to the OpenTracing API.
//
// Tracers implementing this interface should also use the
// SkipContextSetup() function during creation of the span in the
// Start() function to skip the configuration of the context.
type MigrationDeferredContextSetupTracerExtension = migration.DeferredContextSetupTracerExtension

// MigrationOverrideTracerSpanExtension is an interface an OpenTelemetry span
// may implement in order to cooperate with the calls to the
// OpenTracing API.
type MigrationOverrideTracerSpanExtension = migration.OverrideTracerSpanExtension

// WithMigrationDeferredSetup returns a context that can tell the OpenTelemetry
// tracer to skip the context setup in the Start() function.
func WithMigrationDeferredSetup(ctx context.Context) context.Context {
	return migration.WithDeferredSetup(ctx)
}

// SkipMigrationContextSetup can tell the OpenTelemetry tracer to skip the
// context setup during the span creation in the Start() function.
func SkipMigrationContextSetup(ctx context.Context) bool {
	return migration.SkipContextSetup(ctx)
}
