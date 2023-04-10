package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	attemptsPing = 5
	delayPing    = 3 * time.Second
)

func pingDB(db *sql.DB, delay time.Duration, attempts int) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = db.Ping()
		if err == nil {
			return nil
		}
		log.Println("db is not connected, wait...")
		time.Sleep(delay)
	}
	return fmt.Errorf("failed to ping db after %d attempt with %s delay: %w", attempts, delay.String(), err)
}

func NewClientPostgres(cfg StorageConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s search_path=%s",
		cfg.User,
		cfg.DBName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.SSLmode,
		cfg.SearchPath)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("cant parse config: %w", err)
	}
	err = pingDB(db, delayPing, attemptsPing)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	return db, nil
}