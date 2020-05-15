package api

import (
	"errors"
	"net/http"

	"github.com/DataDog/datadog-agent/pkg/process/config"
)

// ErrNotEnabled is a special error type that should be return by a Module factory
// when the given module is simply not enabled.
var ErrNotEnabled = errors.New("module is not enabled")

// Factory encapsulates the initialization of a Module
type Factory struct {
	Name string
	Fn   func(cfg *config.AgentConfig) (Module, error)
}

// Module defines the common API implemented by every System Probe Module
type Module interface {
	GetStats() map[string]interface{}
	Register(*http.ServeMux) error
	Close()
}
