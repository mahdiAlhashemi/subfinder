package subscraping

import (
	"context"
	"net/http"
	"regexp"

	"go.uber.org/ratelimit"
)

// BasicAuth request's Authorization header
type BasicAuth struct {
	Username string
	Password string
}

// Source is an interface inherited by each passive source
type Source interface {
	// Run takes a domain as argument and a session object
	// which contains the extractor for subdomains, http client
	// and other stuff.
	Run(context.Context, string, *Session) <-chan Result
	// Name returns the name of the source.
	Name() string

	// IsDefault returns true if the current source should be
	// used as part of the default execution.
	IsDefault() bool

	// HasRecursiveSupport returns true if the current source
	// accepts subdomains (e.g. subdomain.domain.tld),
	// not just root domains.
	HasRecursiveSupport() bool

	// NeedsKey returns true if the source requires an API key
	NeedsKey() bool

	AddApiKeys([]string)
}

// Session is the option passed to the source, an option is created
// uniquely for each source.
type Session struct {
	// Extractor is the regex for subdomains created for each domain
	Extractor *regexp.Regexp
	// Client is the current http client
	Client *http.Client
	// Rate limit instance
	RateLimiter ratelimit.Limiter
}

// Result is a result structure returned by a source
type Result struct {
	Type   ResultType
	Source string
	Value  string
	Error  error
}

// ResultType is the type of result returned by the source
type ResultType int

// Types of results returned by the source
const (
	Subdomain ResultType = iota
	Error
)
