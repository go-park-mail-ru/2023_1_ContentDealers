package setup

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

func NewClientPostgres() (*sql.DB, error) {
	// TODO: пока хардкод, нужно читать конфиг (cleanenv, viper, gotoenv)
	// Пароли не нужно хранить в конфиге, поэтому нужно думать с переменнными среды
	// 1. https://www.youtube.com/watch?v=sDsAf3gikpQ&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=5
	// 2. https://www.youtube.com/watch?v=asLUNpndGj0&t=1310s

	dsn := "user=postgres dbname=postgres password=postgres host=127.0.0.1 port=5432 sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("cant parse config: %w", err)
	}
	// TODO: нужно несколько раз пробовать подключаться с перерывами
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	return db, nil
}
