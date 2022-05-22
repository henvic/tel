package teljaeger

import (
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/exporters/jaeger"
)

// New returns an OTel Exporter implementation that exports the collected
// spans to Jaeger.
func New(endpointOption EndpointOption) (*Jaeger, error) {
	return jaeger.New(endpointOption)
}

// Jaeger exports OpenTelemetry spans to a Jaeger agent or collector.
type Jaeger = jaeger.Exporter

type EndpointOption = jaeger.EndpointOption

// WithAgentEndpoint configures the Jaeger exporter to send spans to a Jaeger agent
// over compact thrift protocol. This will use the following environment variables for
// configuration if no explicit option is provided:
//
// - OTEL_EXPORTER_JAEGER_AGENT_HOST is used for the agent address host
// - OTEL_EXPORTER_JAEGER_AGENT_PORT is used for the agent address port
//
// The passed options will take precedence over any environment variables and default values
// will be used if neither are provided.
func WithAgentEndpoint(options ...AgentEndpointOption) EndpointOption {
	return jaeger.WithAgentEndpoint(options...)
}

type AgentEndpointOption = jaeger.AgentEndpointOption

// WithAgentHost sets a host to be used in the agent client endpoint.
// This option overrides any value set for the
// OTEL_EXPORTER_JAEGER_AGENT_HOST environment variable.
// If this option is not passed and the env var is not set, "localhost" will be used by default.
func WithAgentHost(host string) AgentEndpointOption {
	return jaeger.WithAgentHost(host)
}

// WithAgentPort sets a port to be used in the agent client endpoint.
// This option overrides any value set for the
// OTEL_EXPORTER_JAEGER_AGENT_PORT environment variable.
// If this option is not passed and the env var is not set, "6831" will be used by default.
func WithAgentPort(port string) AgentEndpointOption {
	return jaeger.WithAgentPort(port)
}

// WithLogger sets a logger to be used by agent client.
func WithLogger(logger *log.Logger) AgentEndpointOption {
	return jaeger.WithLogger(logger)
}

// WithDisableAttemptReconnecting sets option to disable reconnecting udp client.
func WithDisableAttemptReconnecting() AgentEndpointOption {
	return jaeger.WithDisableAttemptReconnecting()
}

// WithAttemptReconnectingInterval sets the interval between attempts to re resolve agent endpoint.
func WithAttemptReconnectingInterval(interval time.Duration) AgentEndpointOption {
	return jaeger.WithAttemptReconnectingInterval(interval)
}

// WithMaxPacketSize sets the maximum UDP packet size for transport to the Jaeger agent.
func WithMaxPacketSize(size int) AgentEndpointOption {
	return jaeger.WithMaxPacketSize(size)
}

// WithCollectorEndpoint defines the full URL to the Jaeger HTTP Thrift collector. This will
// use the following environment variables for configuration if no explicit option is provided:
//
// - OTEL_EXPORTER_JAEGER_ENDPOINT is the HTTP endpoint for sending spans directly to a collector.
// - OTEL_EXPORTER_JAEGER_USER is the username to be sent as authentication to the collector endpoint.
// - OTEL_EXPORTER_JAEGER_PASSWORD is the password to be sent as authentication to the collector endpoint.
//
// The passed options will take precedence over any environment variables.
// If neither values are provided for the endpoint, the default value of "http://localhost:14268/api/traces" will be used.
// If neither values are provided for the username or the password, they will not be set since there is no default.
func WithCollectorEndpoint(options ...CollectorEndpointOption) EndpointOption {
	return jaeger.WithCollectorEndpoint(options...)
}

type CollectorEndpointOption = jaeger.CollectorEndpointOption

// WithEndpoint is the URL for the Jaeger collector that spans are sent to.
// This option overrides any value set for the
// OTEL_EXPORTER_JAEGER_ENDPOINT environment variable.
// If this option is not passed and the environment variable is not set,
// "http://localhost:14268/api/traces" will be used by default.
func WithEndpoint(endpoint string) CollectorEndpointOption {
	return jaeger.WithEndpoint(endpoint)
}

// WithUsername sets the username to be used in the authorization header sent for all requests to the collector.
// This option overrides any value set for the
// OTEL_EXPORTER_JAEGER_USER environment variable.
// If this option is not passed and the environment variable is not set, no username will be set.
func WithUsername(username string) CollectorEndpointOption {
	return jaeger.WithUsername(username)
}

// WithPassword sets the password to be used in the authorization header sent for all requests to the collector.
// This option overrides any value set for the
// OTEL_EXPORTER_JAEGER_PASSWORD environment variable.
// If this option is not passed and the environment variable is not set, no password will be set.
func WithPassword(password string) CollectorEndpointOption {
	return jaeger.WithPassword(password)
}

// WithHTTPClient sets the http client to be used to make request to the collector endpoint.
func WithHTTPClient(client *http.Client) CollectorEndpointOption {
	return jaeger.WithHTTPClient(client)
}
