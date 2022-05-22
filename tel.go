package tel

import (
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// Code is an 32-bit representation of a status state.
type Code = codes.Code

const (
	// Unset is the default status code.
	Unset Code = codes.Unset
	// Error indicates the operation contains an error.
	Error Code = codes.Error
	// OK indicates operation has been validated by an Application developers
	// or Operator to have completed successfully, or contain no error.
	OK Code = codes.Ok
)

// ErrorHandler handles irremediable events.
type ErrorHandler = otel.ErrorHandler

// ErrorHandlerFunc is a convenience adapter to allow the use of a function
// as an ErrorHandler.
type ErrorHandlerFunc = otel.ErrorHandlerFunc

// GetErrorHandler returns the global ErrorHandler instance.
//
// The default ErrorHandler instance returned will log all errors to STDERR
// until an override ErrorHandler is set with SetErrorHandler. All
// ErrorHandler returned prior to this will automatically forward errors to
// the set instance instead of logging.
//
// Subsequent calls to SetErrorHandler after the first will not forward errors
// to the new ErrorHandler for prior returned instances.
func GetErrorHandler() ErrorHandler {
	return otel.GetErrorHandler()
}

// SetErrorHandler sets the global ErrorHandler to h.
//
// The first time this is called all ErrorHandler previously returned from
// GetErrorHandler will send errors to h instead of the default logging
// ErrorHandler. Subsequent calls will set the global ErrorHandler, but not
// delegate errors to h.
func SetErrorHandler(h ErrorHandler) {
	otel.SetErrorHandler(h)
}

// Handle is a convenience function for ErrorHandler().Handle(err).
func Handle(err error) {
	otel.GetErrorHandler().Handle(err)
}

// SetLogger configures the logger used internally to opentelemetry.
func SetLogger(logger logr.Logger) {
	otel.SetLogger(logger)
}

// GetTextMapPropagator returns the global TextMapPropagator. If none has been
// set, a No-Op TextMapPropagator is returned.
func GetTextMapPropagator() TextMapPropagator {
	return otel.GetTextMapPropagator()
}

// SetTextMapPropagator sets propagator as the global TextMapPropagator.
func SetTextMapPropagator(propagator TextMapPropagator) {
	otel.SetTextMapPropagator(propagator)
}

// NewTracer creates a named tracer that implements Tracer interface.
// If the name is an empty string then provider uses default name.
//
// This is short for GetTracerProvider().Tracer(name, opts...)
func NewTracer(name string, opts ...TracerOption) Tracer {
	return otel.Tracer(name, opts...)
}

// GetTracerProvider returns the registered global trace provider.
// If none is registered then an instance of NoopTracerProvider is returned.
//
// Use the trace provider to create a named tracer. E.g.
//     tracer := otel.GetTracerProvider().Tracer("example.com/foo")
// or
//     tracer := otel.Tracer("example.com/foo")
func GetTracerProvider() TracerProvider {
	return otel.GetTracerProvider()
}

// SetTracerProvider registers `tp` as the global trace provider.
func SetTracerProvider(tp TracerProvider) {
	otel.SetTracerProvider(tp)
}

// Version is the current release version of OpenTelemetry in use.
func Version() string {
	return otel.Version()
}
