package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/db/pg"
	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CryptographicAssignment/api/internal/appconfig/iam"
	"github.com/nhan1603/CryptographicAssignment/api/internal/controller/auth"
	"github.com/nhan1603/CryptographicAssignment/api/internal/controller/videos"
	"github.com/nhan1603/CryptographicAssignment/api/internal/model"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository"
	"github.com/nhan1603/CryptographicAssignment/api/internal/repository/generator"
)

func main() {
	log.Println("CryptoGraphy Assignment API")
	ctx := context.Background()

	iamConfig := iam.NewConfig()
	ctx = iam.SetConfigToContext(ctx, iamConfig)

	// Initial DB connection
	conn, err := pg.Connect(os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Initial snowflake generator
	if err := generator.InitSnowflakeGenerators(); err != nil {
		log.Fatal(err)
	}

	// Initial router
	rtr, err := initRouter(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	httpserver.Start(httpserver.Handler(ctx, rtr.routes))
}

func initRouter(
	ctx context.Context,
	db *sql.DB,
) (router, error) {
	repo := repository.New(db)

	return router{
		ctx:       ctx,
		authCtrl:  auth.New(repo, iam.ConfigFromContext(ctx)),
		videoCtrl: videos.New(repo, make(map[*websocket.Conn]bool), make(chan model.NewVideoMessage)),
	}, nil
}
