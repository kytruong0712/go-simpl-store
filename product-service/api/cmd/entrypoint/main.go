package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/kytruong0712/go-simpl-store/product-service/api/cmd/banner"
	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/appconfig/app"
	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/appconfig/db/pg"
	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/appconfig/httpserver"
	sysCtrl "github.com/kytruong0712/go-simpl-store/product-service/api/internal/controller/system"
	sysRest "github.com/kytruong0712/go-simpl-store/product-service/api/internal/handler/rest/system"
	"github.com/kytruong0712/go-simpl-store/product-service/api/internal/repository"
)

func main() {
	banner.Print()
	// Initial config
	config, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	// rootCtx
	ctx := context.Background()

	// Initial DB connection
	conn, err := pg.Connect(config.PGUrl)
	if err != nil {
		log.Fatal("[PG connection error] ", err)
	}
	defer conn.Close()

	// Initial router
	rtr, err := initRouter(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	httpserver.Start(httpserver.Handler(
		httpserver.NewCORSConfig(rtr.corsOrigins),
		sysRest.New(rtr.systemCtrl).CheckDBReady(),
		rtr.routes))
}

func initConfig() (app.Config, error) {
	cfg := app.NewConfig()
	if err := cfg.Validate(); err != nil {
		return app.Config{}, err
	}

	return cfg, nil
}

func initRouter(ctx context.Context, db *sql.DB) (router, error) {
	repo := repository.New(db)

	systemCtrl := sysCtrl.New(repo)

	return router{
		ctx:         ctx,
		systemCtrl:  systemCtrl,
		corsOrigins: []string{"*"},
	}, nil
}
