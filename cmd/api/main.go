package main

import (
	"log"

	"github.com/alejandrogallardo/socialapi/internal/db"
	"github.com/alejandrogallardo/socialapi/internal/env"
	"github.com/alejandrogallardo/socialapi/internal/store"
	"github.com/joho/godotenv"
)

const version = "0.0.1"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", ""),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}
	// Main Database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		// logger.Fatal(err)
		log.Panic(err)
	}

	defer db.Close()
	// logger.Info("database connection pool established")
	log.Println("database connection pool established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
