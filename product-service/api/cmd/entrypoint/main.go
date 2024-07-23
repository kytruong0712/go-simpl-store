package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/kytruong0712/go-simpl-store/product-service/api/cmd/banner"
	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/appconfig/httpserver"
	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/controller/system"
	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/repository"
)

func main() {
	banner.Print()

	// rootCtx
	ctx := context.Background()

	// Initial router
	rtr, err := initRouter(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	httpserver.Start(httpserver.Handler(
		httpserver.NewCORSConfig(rtr.corsOrigins),
		rtr.routes))
}

func initRouter(ctx context.Context, db *sql.DB) (router, error) {
	repo := repository.New(db)

	systemCtrl := system.New(repo)

	return router{
		ctx:         ctx,
		systemCtrl:  systemCtrl,
		corsOrigins: []string{"*"},
	}, nil
}
