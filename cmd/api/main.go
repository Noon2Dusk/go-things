package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/noon2dusk/go-things/db"
	"github.com/noon2dusk/go-things/pkg/api/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	logger := logrus.NewEntry(logrus.StandardLogger())

	queries, err := setupDB()
	if err != nil {
		panic(err.Error())
	}

	apiService := service.New(service.Config{
		Queries: queries,
		Logger:  logger,
	})

	router := apiService.SetupRouter()
	httpErr := make(chan error)
	logger.Info("router started")
	go func() {
		httpErr <- http.ListenAndServe(":80", router)
	}()

	select {
	case err := <-httpErr:
		logger.WithError(err).Fatal("http api server stopped")
	}
}

func setupDB() (*db.Queries, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connecting to db err: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("pinging db err: %w", err)
	}

	query := db.New(conn)

	return query, nil
}
