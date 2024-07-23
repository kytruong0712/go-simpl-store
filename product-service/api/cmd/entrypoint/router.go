package main

import (
	"context"

	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/controller/system"

	"github.com/go-chi/chi/v5"
)

type router struct {
	ctx         context.Context
	corsOrigins []string
	systemCtrl  system.Controller
}

func (rtr router) routes(r chi.Router) {
	// TODO: add routes here
}
