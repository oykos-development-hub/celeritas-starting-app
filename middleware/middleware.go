package middleware

import (
	"myapp/data"

	"github.com/oykos-development-hub/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
