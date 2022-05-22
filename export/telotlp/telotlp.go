package telotlp

import (
	"context"
	"crypto/tls"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/sdk/metric/export/aggregation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// MetricClient manages connections to the collector, handles the
// transformation of data into wire format, and the transmission of that
// data to the collector.
type MetricClient = otlpmetric.Client

// MetricExporter exports metrics data in the OTLP wire format.
type MetricExporter = otlpmetric.Exporter

// NewMetric constructs a new Exporter and starts it.
func NewMetric(ctx context.Context, client MetricClient, opts ...MetricOption) (*MetricExporter, error) {
	return otlpmetric.New(ctx, client, opts...)
}

// NewMetricUnstarted constructs a new Exporter and does not start it.
func NewMetricPUnstarted(client MetricClient, opts ...MetricOption) *MetricExporter {
	return otlpmetric.NewUnstarted(client, opts...)
}

// MetricOption are setting options passed to an Exporter on creation.
type MetricOption = otlpmetric.Option

// WithMetricMetricAggregationTemporalitySelector defines the aggregation.TemporalitySelector used
// for selecting aggregation.Temporality (i.e., Cumulative vs. Delta
// aggregation). If not specified otherwise, exporter will use a
// cumulative temporality selector.
func WithMetricMetricAggregationTemporalitySelector(selector aggregation.TemporalitySelector) MetricOption {
	return otlpmetric.WithMetricAggregationTemporalitySelector(selector)
}

// NewMetricHTTPClient creates a new HTTP metric client.
func NewMetricHTTPClient(opts ...MetricHTTPOption) MetricClient {
	return otlpmetrichttp.NewClient(opts...)
}

// NewMetricHTTP constructs a new Exporter and starts it.
func NewMetricHTTP(ctx context.Context, opts ...MetricHTTPOption) (*MetricExporter, error) {
	return otlpmetrichttp.New(ctx, opts...)
}

// NewMetricHTTPUnstarted constructs a new Exporter and does not start it.
func NewMetricHTTPUnstarted(opts ...MetricHTTPOption) *MetricExporter {
	return otlpmetrichttp.NewUnstarted(opts...)
}

// Compression describes the compression used for payloads sent to the
// collector.
type Compression = otlpmetrichttp.Compression

const (
	// OTLPMetricHTTPNoCompression tells the driver to send payloads without
	// compression.
	MetricHTTPNoCompression = otlpmetrichttp.NoCompression
	// OTLPMetricHTTPGzipCompression tells the driver to send payloads after
	// compressing them with gzip.
	MetricHTTPGzipCompression = otlpmetrichttp.GzipCompression
)

// MetricHTTPOption applies an option to the HTTP client.
type MetricHTTPOption = otlpmetrichttp.Option

// MetricHTTPRetryConfig defines configuration for retrying batches in case of export
// failure using an exponential backoff.
type MetricHTTPRetryConfig = otlpmetrichttp.RetryConfig

// WithMetricHTTPEndpoint allows one to set the address of the collector endpoint that
// the driver will use to send metrics. If unset, it will instead try to use
// the default endpoint (localhost:4318). Note that the endpoint must not
// contain any URL path.
func WithMetricHTTPEndpoint(endpoint string) MetricHTTPOption {
	return otlpmetrichttp.WithEndpoint(endpoint)
}

// WithMetricHTTPCompression tells the driver to compress the sent data.
func WithMetricHTTPCompression(compression Compression) MetricHTTPOption {
	return otlpmetrichttp.WithCompression(compression)
}

// WithMetricHTTPURLPath allows one to override the default URL path used
// for sending metrics. If unset, default ("/v1/metrics") will be used.
func WithMetricHTTPURLPath(urlPath string) MetricHTTPOption {
	return otlpmetrichttp.WithURLPath(urlPath)
}

// WithMetricHTTPMaxAttempts allows one to override how many times the driver
// will try to send the payload in case of retryable errors.
// The max attempts is limited to at most 5 retries. If unset,
// default (5) will be used.
//
// Deprecated: Use WithRetry instead.
func WithMetricHTTPMaxAttempts(maxAttempts int) MetricHTTPOption {
	return otlpmetrichttp.WithMaxAttempts(maxAttempts)
}

// WithMetricHTTPBackoff tells the driver to use the duration as a base of the
// exponential backoff strategy. If unset, default (300ms) will be
// used.
//
// Deprecated: Use WithRetry instead.
func WithMetricHTTPBackoff(duration time.Duration) MetricHTTPOption {
	return otlpmetrichttp.WithBackoff(duration)
}

// WithMetricHTTPTLSClientConfig can be used to set up a custom TLS
// configuration for the client used to send payloads to the
// collector. Use it if you want to use a custom certificate.
func WithMetricHTTPTLSClientConfig(tlsCfg *tls.Config) MetricHTTPOption {
	return otlpmetrichttp.WithTLSClientConfig(tlsCfg)
}

// WithMetricHTTPInsecure tells the driver to connect to the collector using the
// HTTP scheme, instead of HTTPS.
func WithMetricHTTPInsecure() MetricHTTPOption {
	return otlpmetrichttp.WithInsecure()
}

// WithMetricHTTPHeaders allows one to tell the driver to send additional HTTP
// headers with the payloads. Specifying headers like Content-Length,
// Content-Encoding and Content-Type may result in a broken driver.
func WithMetricHTTPHeaders(headers map[string]string) MetricHTTPOption {
	return otlpmetrichttp.WithHeaders(headers)
}

// WithMetricHTTPTimeout tells the driver the max waiting time for the backend to process
// each metrics batch.  If unset, the default will be 10 seconds.
func WithMetricHTTPTimeout(duration time.Duration) MetricHTTPOption {
	return otlpmetrichttp.WithTimeout(duration)
}

// WithMetricHTTPRetry configures the retry policy for transient errors that may occurs
// when exporting traces. An exponential back-off algorithm is used to ensure
// endpoints are not overwhelmed with retries. If unset, the default retry
// policy will retry after 5 seconds and increase exponentially after each
// error for a total of 1 minute.
func WithMetricHTTPRetry(rc MetricHTTPRetryConfig) MetricHTTPOption {
	return otlpmetrichttp.WithRetry(rc)
}

// NewOTLPGRPCMetricClient creates a new gRPC metric client.
func NewOTLPGRPCMetricClient(opts ...GRPCOption) MetricClient {
	return otlpmetricgrpc.NewClient()
}

// NewOTLPGRPCMetric constructs a new Exporter and starts it.
func NewOTLPGRPCMetric(ctx context.Context, opts ...GRPCOption) (*MetricExporter, error) {
	return otlpmetric.New(ctx, NewOTLPGRPCMetricClient(opts...))
}

// NewOTLPGPRCetricUnstarted constructs a new Exporter and does not start it.
func NewOTLPGRPCMetricUnstarted(opts ...GRPCOption) *MetricExporter {
	return otlpmetric.NewUnstarted(NewOTLPGRPCMetricClient(opts...))
}

// GRPCOption applies an option to the gRPC driver.
type GRPCOption = otlpmetricgrpc.Option

// OTLPGRPCRetryConfig defines configuration for retrying export of span batches that
// failed to be received by the target endpoint.
//
// This configuration does not define any network retry strategy. That is
// entirely handled by the gRPC ClientConn.
type OTLPGRPCRetryConfig = otlpmetricgrpc.RetryConfig

// WithInsecure disables client transport security for the exporter's gRPC
// connection just like grpc.WithInsecure()
// (https://pkg.go.dev/google.golang.org/grpc#WithInsecure) does. Note, by
// default, client security is required unless WithInsecure is used.
//
// This option has no effect if WithGRPCConn is used.
func WithGRPCInsecure() GRPCOption {
	return otlpmetricgrpc.WithInsecure()
}

// WithEndpoint sets the target endpoint the exporter will connect to. If
// unset, localhost:4317 will be used as a default.
//
// This option has no effect if WithGRPCConn is used.
func WithGRPCEndpoint(endpoint string) GRPCOption {
	return otlpmetricgrpc.WithEndpoint(endpoint)
}

// WithGRPCReconnectionPeriod set the minimum amount of time between connection
// attempts to the target endpoint.
//
// This option has no effect if WithGRPCConn is used.
func WithGRPCReconnectionPeriod(rp time.Duration) GRPCOption {
	return otlpmetricgrpc.WithReconnectionPeriod(rp)
}

// WithGRPCCompressor sets the compressor for the gRPC client to use when sending
// requests. It is the responsibility of the caller to ensure that the
// compressor set has been registered with google.golang.org/grpc/encoding.
// This can be done by encoding.RegisterCompressor. Some compressors
// auto-register on import, such as gzip, which can be registered by calling
// `import _ "google.golang.org/grpc/encoding/gzip"`.
//
// This option has no effect if WithGRPCConn is used.
func WithGRPCCompressor(compressor string) GRPCOption {
	return otlpmetricgrpc.WithCompressor(compressor)
}

// WithGRPCHeaders will send the provided headers with each gRPC requests.
func WithGRPCHeaders(headers map[string]string) GRPCOption {
	return otlpmetricgrpc.WithHeaders(headers)
}

// WithGRPCTLSCredentials allows the connection to use TLS credentials when
// talking to the server. It takes in grpc.TransportCredentials instead of say
// a Certificate file or a tls.Certificate, because the retrieving of these
// credentials can be done in many ways e.g. plain file, in code tls.Config or
// by certificate rotation, so it is up to the caller to decide what to use.
//
// This option has no effect if WithGRPCConn is used.
func WithGRPCTLSCredentials(creds credentials.TransportCredentials) GRPCOption {
	return otlpmetricgrpc.WithTLSCredentials(creds)
}

// WithGRPCServiceConfig defines the default gRPC service config used.
//
// This option has no effect if WithGRPCConn is used.
func WithGRPCServiceConfig(serviceConfig string) GRPCOption {
	return otlpmetricgrpc.WithServiceConfig(serviceConfig)
}

// WithGRPCDialOption sets explicit grpc.DialOptions to use when making a
// connection. The options here are appended to the internal grpc.DialOptions
// used so they will take precedence over any other internal grpc.DialOptions
// they might conflict with.
//
// This option has no effect if WithGRPCConn is used.
func WithGRPCDialOption(opts ...grpc.DialOption) GRPCOption {
	return otlpmetricgrpc.WithDialOption(opts...)
}

// WithGRPCConn sets conn as the gRPC ClientConn used for all communication.
//
// This option takes precedence over any other option that relates to
// establishing or persisting a gRPC connection to a target endpoint. Any
// other option of those types passed will be ignored.
//
// It is the callers responsibility to close the passed conn. The client
// Shutdown method will not close this connection.
func WithGRPCConn(conn *grpc.ClientConn) GRPCOption {
	return otlpmetricgrpc.WithGRPCConn(conn)
}

// WithGRPCTimeout sets the max amount of time a client will attempt to export a
// batch of spans. This takes precedence over any retry settings defined with
// WithRetry, once this time limit has been reached the export is abandoned
// and the batch of spans is dropped.
//
// If unset, the default timeout will be set to 10 seconds.
func WithGRPCTimeout(duration time.Duration) GRPCOption {
	return otlpmetricgrpc.WithTimeout(duration)
}

// WithGRPCRetry sets the retry policy for transient retryable errors that may be
// returned by the target endpoint when exporting a batch of spans.
//
// If the target endpoint responds with not only a retryable error, but
// explicitly returns a backoff time in the response. That time will take
// precedence over these settings.
//
// These settings do not define any network retry strategy. That is entirely
// handled by the gRPC ClientConn.
//
// If unset, the default retry policy will be used. It will retry the export
// 5 seconds after receiving a retryable error and increase exponentially
// after each error for no more than a total time of 1 minute.
func WithGRPCRetry(settings OTLPGRPCRetryConfig) GRPCOption {
	return otlpmetricgrpc.WithRetry(settings)
}
