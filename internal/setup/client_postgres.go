package setup

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	maxOpenConns = 10
)

func NewClientPostgres(cfg StorageConfig) (*sql.DB, error) {
	// TODO: пока хардкод, нужно читать конфиг (cleanenv, viper, gotoenv)
	// Пароли не нужно хранить в конфиге, поэтому нужно думать с переменнными среды
	// 1. https://www.youtube.com/watch?v=sDsAf3gikpQ&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=5
	// 2. https://www.youtube.com/watch?v=asLUNpndGj0&t=1310s

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s search_path=filmium",
		cfg.User,
		cfg.DBName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.SSLmode)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("cant parse config: %w", err)
	}
	// TODO: нужно несколько раз пробовать подключаться с перерывами
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConns)
	return db, nil
}
