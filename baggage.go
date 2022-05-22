package tel

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
)

// BaggageProperty is an additional metadata entry for a baggage list-member.
type BaggageProperty = baggage.Property

func NewKeyProperty(key string) (BaggageProperty, error) {
	return baggage.NewKeyProperty(key)
}

func NewKeyValueProperty(key, value string) (BaggageProperty, error) {
	return baggage.NewKeyValueProperty(key, value)
}

// BaggageMember is a list-member of a baggage-string as defined by the W3C Baggage
// specification.
type BaggageMember = baggage.Member

// NewMember returns a new Member from the passed arguments. An error is
// returned if the created Member would be invalid according to the W3C
// Baggage specification.
func NewMember(key, value string, props ...BaggageProperty) (BaggageMember, error) {
	return baggage.NewMember(key, value, props...)
}

// Baggage is a list of baggage members representing the baggage-string as
// defined by the W3C Baggage specification.
type Baggage = baggage.Baggage

// NewBaggage returns a new valid Baggage. It returns an error if it results in a
// Baggage exceeding limits set in that specification.
//
// It expects all the provided members to have already been validated.
func NewBaggage(members ...BaggageMember) (Baggage, error) {
	return baggage.New(members...)
}

// ParseBaggage attempts to decode a baggage-string from the passed string. It
// returns an error if the input is invalid according to the W3C Baggage
// specification.
//
// If there are duplicate list-members contained in baggage, the last one
// defined (reading left-to-right) will be the only one kept. This diverges
// from the W3C Baggage specification which allows duplicate list-members, but
// conforms to the OpenTelemetry Baggage specification.
func ParseBaggage(bStr string) (Baggage, error) {
	return baggage.Parse(bStr)
}

// ContextWithBaggage returns a copy of parent with baggage.
func ContextWithBaggage(parent context.Context, b Baggage) context.Context {
	// Delegate so any hooks for the OpenTracing bridge are handled.
	return baggage.ContextWithBaggage(parent, b)
}

// ContextWithoutBaggage returns a copy of parent with no baggage.
func ContextWithoutBaggage(parent context.Context) context.Context {
	// Delegate so any hooks for the OpenTracing bridge are handled.
	return baggage.ContextWithoutBaggage(parent)
}

// FromContext returns the baggage contained in ctx.
func FromContext(ctx context.Context) Baggage {
	// Delegate so any hooks for the OpenTracing bridge are handled.
	return baggage.FromContext(ctx)
}
