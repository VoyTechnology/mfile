// Package mfile allows to easily read data from different sources just by
// specifying a schema, and importing a provider which implements that schema.
package mfile // import "github.com/voytechnology/mfile"

import (
	"net/url"
)

// Error defines the errors returned from the mfile package
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrParsing is returned when its impossible to parse the given path
	ErrParsing = Error("mfile: unable to parse path")
	// ErrSchemaExists is returned when its impossible to register a schema,
	// because one already exists.
	ErrSchemaExists = Error("mfile: schema already exists")
	// ErrUnknownSchema is returned when the user attempted to use a schema
	// which was not registered
	ErrUnknownSchema = Error("mfile: unknown schema")
)

// Handler is the main schema that each handler must implement in order to be
// used.
type Handler interface {
	// ReadFile reads the full file and returns the result.
	ReadFile(path string) ([]byte, error)
}

var schemas = make(map[string]Handler)

func load(path string) (Handler, string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, "", ErrParsing
	}

	handler, exists := schemas[u.Scheme]
	if !exists {
		return nil, "", ErrUnknownSchema
	}

	// clear the schema as we no longer care about it.
	u.Scheme = ""

	return handler, u.String(), nil
}

// Register is used by schema providers to specify the handlers for a schema of
// a given name.
func Register(name string, handler Handler) error {
	if _, exists := schemas[name]; exists {
		return ErrSchemaExists
	}

	schemas[name] = handler
	return nil
}
