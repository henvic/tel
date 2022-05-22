# tel: Open-Telemetry wrapper API for Go
[![Go Reference](https://pkg.go.dev/badge/github.com/henvic/tel.svg)](https://pkg.go.dev/github.com/henvic/tel)

This project is a wrapper for the [OpenTelemetry Go API and SDK](github.com/open-telemetry/opentelemetry-go/).

## Goals
* Make it possible to use the OpenTelemetry API without having to part ways with the idioms of the GO programming language.
* Create a more concise API that is easier to use and discover.
* Encourage changing the go.opentelemetry.io API and SDK.

## Why
The current API for go.opentelemetry.io is hard to use as it exposes dozens of packages, requiring someone to be very familiar with OpenTelemetry before taking advantage of its APIs.

This makes it very uninviting for Go developers who are not telemetry experts to instrument their code. Furthermore, many package names used are common names, leading to conflicts with function/variable names or other packages – which can be a problem, especially when relying on auto-completion.

For example: a package called go.opentelemetry.io/sdk/metric/controller/time, is exported as time, even though the standard library has a package named time.

It is important to remember that there is no concept of subpackage in Go and that a package name should speak clearly about what it is, independently of its whole path.

> Packages net/http/httptest or net/http/httptrace contain the prefix http exactly to mitigate this problem.

## Status
1. The code here was mapped manually, with minor code editor tooling helping.
2. Code generation was not considered for this prototype but would be essential to move things forward.
3. go.opentelemetry.io/semconv was not included.
## semconv
Package semconv implements OpenTelemetry semantic conventions. OpenTelemetry semantic conventions are agreed standardized naming patterns for OpenTelemetry things.

In the [opentelemetry-go/semconv](https://github.com/open-telemetry/opentelemetry-go/tree/main/semconv) directory in the official repository you're going to find the following:

```sh
├── internal
│   ├── http.go
│   └── http_test.go
├── template.j2
├── v1.10.0
│   └── doc.go, exception.go, http.go, resource.go, schema.go, trace.go
├── v1.4.0
│   └── doc.go, exception.go, http.go, resource.go, schema.go, trace.go
├── v1.5.0
│   └── doc.go, exception.go, http.go, resource.go, schema.go, trace.go
├── v1.6.1
│   └── doc.go, exception.go, http.go, resource.go, schema.go, trace.go
├── v1.7.0
│   └── doc.go, exception.go, http.go, resource.go, schema.go, trace.go
├── v1.8.0
│   └── doc.go, exception.go, http.go, resource.go, schema.go, trace.go
└── v1.9.0
    └── doc.go, exception.go, http.go, resource.go, schema.go, trace.go
    
    8 directories, 45 files
```

For each version, there will be a set of Go files there.
Instead, there could be a separate repository, perhaps called gosemconv, tagged for each version. This would mean just quickly checking go.mod would be enough for most people to define what to use.

In case someone requires support for multiple versions, they could resort to the following [Go proverb](https://go-proverbs.github.io/) and solve it themselves:

> A little copying is better than a little dependency.

This is a better trade-off than versioning each different version in directories. It will make it less likely for a majority to use an old version by mistake and negatively impact only a few users.

## Future
If this proposal:

1. Succeeds in improving the official API: success; archive this.
2. Fails: API code generation mapping the official API. This is a suboptimal outcome, but might be worth.
