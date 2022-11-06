package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/noon2dusk/go-things/db"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	logger := logrus.NewEntry(logrus.StandardLogger())
	conn, query, err := setupDB()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	runMigrations(conn, query, logger, "./database/migration/")
}

func setupDB() (*sql.DB, *db.Queries, error) {
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
		return nil, nil, fmt.Errorf("connecting to db err: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, nil, fmt.Errorf("pinging db err: %w", err)
	}

	query := db.New(conn)

	return conn, query, nil
}

func runMigrations(conn *sql.DB, query *db.Queries, logger *logrus.Entry, dbDir string) {
	_, err := conn.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s;`, os.Getenv("MYSQL_DATABASE")))

	counter := struct {
		skipped  int
		migrated int
	}{}
	ctx := context.Background()

	fileInfo, err := ioutil.ReadDir(dbDir)
	if err != nil {
		logger.WithError(err).Fatal("Error reading DB migration dir")
	}

	_, err = conn.Exec(`CREATE TABLE IF NOT EXISTS migration
		(
			id INT NOT NULL AUTO_INCREMENT,
			executed_at TIMESTAMP default now(),
			name VARCHAR(255) NOT NULL,
			PRIMARY KEY (id)
		);`)
	if err != nil {
		logger.WithError(err).Fatal("could not create migration table")
	}

	for _, file := range fileInfo {
		log := logger.WithField("file", file.Name())

		content, err := ioutil.ReadFile(dbDir + file.Name())
		if err != nil {
			log.WithError(err).Fatal("could not read DB migration file")
		}

		_, err = query.GetMigrationByName(ctx, file.Name())
		if err != nil {
			if !errors.Is(sql.ErrNoRows, err) {
				log.WithError(err).Fatal("could not get migration from DB")
			}
		} else {
			continue
		}

		start := time.Now()
		_, err = conn.Exec(string(content))
		if err != nil {
			log.WithError(err).Fatal("failed executing DB migration file")
		}
		end := time.Since(start)

		err = query.AddMigration(ctx, file.Name())
		if err != nil {
			log.WithError(err).Fatal("could not write migration to migration table")
		}

		log.WithField("t", end).Info("successfully executed migration")
		counter.migrated++
	}

	logger.WithFields(logrus.Fields{
		"dir":      dbDir,
		"skipped":  counter.skipped,
		"migrated": counter.migrated,
	}).Info("migration finished")
}
