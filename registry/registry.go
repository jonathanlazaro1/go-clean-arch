package registry

import (
	"database/sql"

	controller "github.com/jonathanlazaro1/go-clean-arch/interface/controller"
)

type registry struct {
	db *sql.DB
}

// Registry has contracts to register dependencies
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry creates a new instance of registry
func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewAppController()
}
