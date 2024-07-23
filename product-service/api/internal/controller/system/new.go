package system

import (
	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
